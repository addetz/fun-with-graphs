package destination_city


func destCity(paths [][]string) string {
    nodes := make(map[string]string, 0)
    for _, v := range paths {
        nodes[v[0]]= v[1]
    }
    for _, end := range nodes {
        // there is no outgoing for the end destination
        if _, ok := nodes[end]; !ok {
            return end
        }
    }
    return ""
}
