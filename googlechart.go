package googlecharts

type GoogleChartColumn struct {
	Id    string `json:"id"`
	Label string `json:"label"`
	Type  string `json:"type"`
}

type GoogleChartRow struct {
	Cells []GoogleChartCell `json:"c"`
}

func (row *GoogleChartRow) AddCell(value interface{}, tooltip string) {
	row.Cells = append(row.Cells, GoogleChartCell{Value: value, Tooltip: tooltip})
}

type GoogleChartCell struct {
	Value   interface{} `json:"v"`
	Tooltip string      `json:"p"`
}

type GoogleChartDataTable struct {
	Columns []GoogleChartColumn `json:"cols"`
	Rows    []*GoogleChartRow   `json:"rows"`
}

func (table *GoogleChartDataTable) AddColumn(id string, columntype string, label string) {
	table.Columns = append(table.Columns, GoogleChartColumn{id, label, columntype})
}

func NewGoogleChartDataTable() *GoogleChartDataTable {

	dt := &GoogleChartDataTable{}
	dt.Columns = make([]GoogleChartColumn, 0)
	dt.Rows = make([]*GoogleChartRow, 0)

	return dt

}

func (table *GoogleChartDataTable) AddRow() *GoogleChartRow {
	row := &GoogleChartRow{}
	row.Cells = make([]GoogleChartCell, 0)
	table.Rows = append(table.Rows, row)
	return row
}
