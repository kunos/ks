package khtml

type Table struct {
	Item
	rows  int
	cols  int
	cells []*TableCell
}

func (this *Table) Render(code *string) {

	*code = *code + `<table>`

	this.Item.Render(code)

	*code = *code + `</table>`

}

func NewTable(rows int, columns int) *Table {
	nt := Table{rows: rows, cols: columns}

	// ADD THE rows
	for i := 0; i < rows; i++ {

		nr := TableRow{}

		nt.AddElement(&nr)

		for c := 0; c < columns; c++ {
			nc := TableCell{}
			nr.AddElement(&nc)
			nt.cells = append(nt.cells, &nc)
		}

	}

	return &nt
}

func (this *Table) GetCellAt(row int, col int) *TableCell {
	return this.cells[row*this.cols+col]
}

// TABLE ROW
type TableRow struct {
	Item
}

func (this *TableRow) Render(code *string) {
	*code = *code + `<tr>`

	this.Item.Render(code)

	*code = *code + `</tr>`

}

// COLUMN

type TableCell struct {
	Item
	Text string
}

func (this *TableCell) Render(code *string) {
	*code = *code + `<td>` + this.Text

	this.Item.Render(code)

	*code = *code + `</td>`

}
