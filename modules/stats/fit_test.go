package stats

import (
	"testing"
)

func TestPerfectLeastSquaresFit(t *testing.T) {
	records := [][]float64{
		{1, 3}, {2, 5}, {3, 7}, {4, 9}, {5, 11}, {6, 13}, {7, 15},
	}

	m, dm, b, db, rSquared := LeastSquaresFit(records, 0, 1)

	if m != 2 || dm != 0 || b != 1 || db != 0 || rSquared != 1 {
		t.Errorf("LeastSquaresFit returned unexpected values: m=%f, dm=%f, b=%f, db=%f, rSquared=%f", m, dm, b, db, rSquared)
	}
}

func TestImperfectLeastSquaresFit(t *testing.T) {
	records := [][]float64{
		{1, 2.85}, {2, 5.10}, {3, 6.9}, {4, 9.1}, {5, 10.9}, {6, 12.85}, {7, 15.2},
	}

	m, dm, b, db, rSquared := LeastSquaresFit(records, 0, 1)

	deviance := 1.0e-6

	if m-2.019643 > deviance ||
		dm-0.056762 > deviance ||
		b-0.907143 > deviance ||
		db-0.253848 > deviance ||
		rSquared-0.999014 > deviance {
		t.Errorf("LeastSquaresFit returned unexpected values: m=%f, dm=%f, b=%f, db=%f, rSquared=%f", m, dm, b, db, rSquared)
	}
}
