package mocks

import "time-tracker-be/internal/models"

type RecordStoreMock struct {
	idGen   int
	records []*models.Record
}

func (m *RecordStoreMock) Add(r *models.Record) (*models.Record, error)  {
	m.idGen += 1
	rec := &models.Record{
		ID:       m.idGen,
		Name:     r.Name,
		Start:    r.Start,
		Finish:   r.Finish,
		Duration: r.Finish - r.Start,
	}
	m.records = append(m.records, rec)
	return rec, nil
}

func (m *RecordStoreMock) List(from int64) ([]*models.Record, error)  {
	return m.records, nil
}

func NewRecordStoreMock(idGen int, recs []*models.Record) *RecordStoreMock {
	return &RecordStoreMock{idGen:idGen, records:recs}
}