{{ range . }}
type {{.StructName}}ID {{.IDType}}
type {{.StructName}}Slice []*{{.StructName}}
// GroupByID returns a map and a slice of given ItemSlice
func (ss {{.StructName}}Slice) GroupByID() (grouped map[{{.StructName}}ID]*{{.StructName}}, ids []{{.StructName}}ID){
    if len(ss) == 0 {
        return
    }
    grouped = make(map[{{.StructName}}ID]*{{.StructName}})
    ids = make([]{{.StructName}}ID, len(ss))
    for idx, i := range ss {
        grouped[i.Id] = i
        ids[idx] = {{.StructName}}ID(i.Id)
    }
    return
}
{{ end }}
