package sql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecutorTrain(t *testing.T) {
	a := assert.New(t)
	a.NotPanics(func() {
		e := run(testTrainSelectIris, testCfg)
		a.NoError(e)
	})
}

func TestExecutorInfer(t *testing.T) {
	a := assert.New(t)
	a.NotPanics(func() {
		e := run(testPredictSelectIris, testCfg)
		a.EqualError(e, "infer not implemented")
	})
}

func TestCreatePredictionTable(t *testing.T) {
	a := assert.New(t)
	trainParsed, e := newParser().Parse(testTrainSelectIris)
	a.NoError(e)
	inferParsed, e := newParser().Parse(testPredictSelectIris)
	a.NoError(e)
	a.NoError(createPredictionTable(trainParsed, inferParsed, testCfg))
}