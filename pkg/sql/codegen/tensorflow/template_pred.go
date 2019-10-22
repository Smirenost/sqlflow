// Copyright 2019 The SQLFlow Authors. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tensorflow

import "sqlflow.org/sqlflow/pkg/sql/codegen"

type predFiller struct {
	DataSource  string
	Select      string
	ResultTable string
	// below members comes from TrainIR
	Estimator         string
	IsKerasModel      bool
	FieldMetas        []*codegen.FieldMeta
	FeatureColumnCode string
	Y                 *codegen.FieldMeta
	ModelParams       map[string]interface{}
	Save              string
}

const tfPredTemplateText = `
from sqlflow_submitter.tensorflow.predict import pred

feature_column_names = [{{range .FieldMetas}}
"{{.Name}}",
{{end}}]

feature_metas = dict()
{{ range $value := .FieldMetas }}
feature_metas["{{$value.Name}}"] = {
    "feature_name": "{{$value.Name}}",
    "dtype": "{{$value.DType | dtypeToString}}",
    "delimiter": "{{$value.Delimiter}}",
    "shape": {{$value.Shape | intArrayToJSONString}},
    "is_sparse": "{{$value.IsSparse}}" == "true"
}
{{end}}

label_meta = {
    "feature_name": "{{.Y.Name}}",
    "dtype": "{{.Y.DType | dtypeToString}}",
    "delimiter": "{{.Y.Delimiter}}",
    "shape": {{.Y.Shape | intArrayToJSONString}},
    "is_sparse": "{{.Y.IsSparse}}" == "true"
}

model_params=dict()
{{range $k, $v := .ModelParams}}
model_params["{{$k}}"]={{$v | attrToPythonValue}}
{{end}}

pred(is_keras_model="{{.IsKerasModel}}" == "true",
    datasource="{{.DataSource}}",
    estimator="{{.Estimator}}",
    select="""{{.Select}}""",
    result_table="{{.ResultTable}}",
    feature_column_code="""{{.FeatureColumnCode}}""",
    feature_column_names=feature_column_names,
    feature_metas=feature_metas,
    label_meta=label_meta,
    model_params=model_params,
    save="{{.Save}}",
    batch_size=1)
`