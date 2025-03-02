package upsert

import (
	"errors"
	"fmt"
	"strings"
)

const (
	baseQuery = `
		MERGE (u:User {name: $user.name})
		ON CREATE SET %s
		ON MATCH SET %s

		MERGE (h:Host {id: $host.id})
		ON CREATE SET %s
		ON MATCH SET %s

		MERGE (u)-[r1:LOGGED_ON]->(h)
`
	processQuery = `
		MERGE (%s:Process {id: $%s.id})
		ON CREATE SET %s
		ON MATCH SET %s

		MERGE (u)-[r%s:INITIATED {created_at: $createdAt}]->(%s)
		MERGE (%s)-[r2%s:RUN_ON {created_at: $createdAt}]->(h)
`
)

func queryBuilder(params map[string]any) (string, error) {
	b := strings.Builder{}

	u, ok := params[mapKeyUser].(map[string]any)
	if !ok {
		return "", errors.New("user is not a map[string]map[string]any")
	}
	varsUser := keysToVarList("user", "u", GetKeys(u))

	h, ok := params[mapKeyHost].(map[string]any)
	if !ok {
		return "", errors.New("host is not a map[string]map[string]any")
	}

	varsHost := keysToVarList("host", "h", GetKeys(h))
	q := fmt.Sprintf(baseQuery, varsUser, varsUser, varsHost, varsHost)

	b.WriteString(q)

	p, ok := params[mapKeyProcesses].(map[string]any)
	if !ok {
		return "", errors.New("processes is not a map[string]map[string]any")
	}

	for i, pMap := range p {
		list, ok := pMap.(map[string]any)
		if !ok {
			return "", errors.New("processes is not a map[string]map[string]any")
		}
		ref := fmt.Sprintf("%s.%s", mapKeyProcesses, i)
		vars := keysToVarList(ref, i, GetKeys(list))
		partial := fmt.Sprintf(processQuery, i, ref, vars, vars, i, i, i, i)
		b.WriteString(partial)
	}
	return b.String(), nil

}

func GetKeys(m map[string]any) []string {
	keys := make([]string, 0)
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func keysToVarList(pefix, nodeReference string, keys []string) string {
	varList := make([]string, len(keys))
	for idx, key := range keys {
		varList[idx] = fmt.Sprintf("%s.%s = $%s.%s", nodeReference, key, pefix, key)
	}
	return strings.Join(varList, ", ")
}
