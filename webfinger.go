package main

func HandleRel(j JRD, rel string) JRD {
	if rel != "" {
		currentLinks := j.Links
		if currentLinks != nil {
			links := []Link{}
			for _, link := range currentLinks {
				if link.Rel == rel {
					links = append(links, link)
				}
			}
			j.Links = links
		}
	}
	return j
}
