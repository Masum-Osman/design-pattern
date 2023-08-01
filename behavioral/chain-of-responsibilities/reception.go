package chainofresponsibilities

type Reception struct {
	next Department
}

func (r *Reception) execute(p *Patient) {

}
