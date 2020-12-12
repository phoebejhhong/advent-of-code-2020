package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type bagNode struct {
	childRules []childRule
	parents    []string
}
type childRule struct {
	name  string
	count int64
}

func parseContent(content string) []childRule {
	childRules := []childRule{}
	if content == "no other bags" {
		return childRules
	}

	contentRegex := regexp.MustCompile(`(\d)\s([a-z]+\s[a-z]+)\sbag`)
	result := contentRegex.FindAllStringSubmatch(content, -1)
	for _, individualResult := range result {
		count, _ := strconv.ParseInt(individualResult[1], 10, 64)
		name := individualResult[2]
		childRules = append(childRules, childRule{name, count})
	}

	return childRules
}

func parseRule(rule string) (string, []childRule) {
	mainRegex := regexp.MustCompile(`^(.+\s.+)\sbags\scontain\s(.+)\.$`)
	result := mainRegex.FindStringSubmatch(rule)
	bagName := result[1]
	content := result[2]
	childRules := parseContent(content)

	return bagName, childRules
}

func parseRules(rules []string) map[string]*bagNode {
	bags := map[string]*bagNode{}
	for _, rule := range rules {
		if rule != "" {
			bagName, content := parseRule(rule)
			if bags[bagName] == nil {
				bags[bagName] = &bagNode{}
			}
			bag := bags[bagName]
			for _, contentRule := range content {
				contentName := contentRule.name
				if bags[contentName] == nil {
					bags[contentName] = &bagNode{}
				}
				contentBag := bags[contentName]

				bag.childRules = append(bag.childRules, childRule{name: contentName, count: contentRule.count})
				contentBag.parents = append(contentBag.parents, bagName)
			}
		}
	}
	return bags
}

func getUniqueParentMapFor(bagName string, bags map[string]*bagNode) map[string]bool {
	parentMap := map[string]bool{}
	bag := bags[bagName]
	for _, parentName := range bag.parents {
		parentMap[parentName] = true
		upperParentsMap := getUniqueParentMapFor(parentName, bags)
		for upperParent, _ := range upperParentsMap {
			parentMap[upperParent] = true
		}
	}
	return parentMap
}

func getAllChildCountFor(bagName string, bags map[string]*bagNode) int64 {
	bag := bags[bagName]
	var childCount int64 = 0
	for _, childRule := range bag.childRules {
		childCount += childRule.count * (getAllChildCountFor(childRule.name, bags) + 1)
	}
	return childCount
}

func main() {
	data, _ := ioutil.ReadFile("./07-input.txt")
	rules := strings.Split(string(data), "\n")
	bags := parseRules(rules)
	parentMap := getUniqueParentMapFor("shiny gold", bags)
	childrenCount := getAllChildCountFor("shiny gold", bags)
	fmt.Println(" 📝 Possible outer bag for shiny gold is", len(parentMap))
	fmt.Println(" 📝 Required number of bags inside shiny gold bag is", childrenCount)
}
