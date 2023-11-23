func entityParser(text string) string {
    entityMap := map[string]string{
        "&quot;": "\"",
        "&apos;": "'",
        "&gt;": ">",
        "&lt;": "<",
        "&frasl;": "/",
        "&amp;": "&",
    }

    i, n := 0, len(text)
    res := make([]string, 0)
    for i < n{
        isEntity := false
        if text[i] == '&'{
            for k, v := range entityMap{
                if i + len(k) <= n && text[i: i+len(k)] == k{
                    res = append(res, v)
                    isEntity = true
                    i += len(k)
                    break
                }
            }
        }
        if !isEntity{
            res = append(res, text[i:i+1])
            i++
        }
    }
    return strings.Join(res, "")
}