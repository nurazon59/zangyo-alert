package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Prompter struct {
	r *bufio.Reader
	w io.Writer
}

func NewPrompterFrom(r io.Reader, w io.Writer) *Prompter {
	return &Prompter{r: bufio.NewReader(r), w: w}
}

func (p *Prompter) AskYesNo(prompt string) (bool, error) {
	for {
		fmt.Fprintf(p.w, "%s ", prompt)
		line, err := p.readLine()
		if err != nil {
			return false, err
		}
		switch strings.ToLower(strings.TrimSpace(line)) {
		case "y", "yes":
			return true, nil
		case "n", "no":
			return false, nil
		}
	}
}

func (p *Prompter) AskInt(prompt string, def *int) (int, error) {
	for {
		if def != nil {
			fmt.Fprintf(p.w, "%s [%d] ", prompt, *def)
		} else {
			fmt.Fprintf(p.w, "%s ", prompt)
		}
		line, err := p.readLine()
		if err != nil {
			return 0, err
		}
		s := strings.TrimSpace(line)
		if s == "" && def != nil {
			return *def, nil
		}
		v, err := strconv.Atoi(strings.ReplaceAll(s, ",", ""))
		if err == nil {
			return v, nil
		}
		fmt.Fprintf(p.w, "整数として解釈できません: %q\n", s)
	}
}

func (p *Prompter) AskFloat(prompt string, def *float64) (float64, error) {
	for {
		if def != nil {
			fmt.Fprintf(p.w, "%s [%g] ", prompt, *def)
		} else {
			fmt.Fprintf(p.w, "%s ", prompt)
		}
		line, err := p.readLine()
		if err != nil {
			return 0, err
		}
		s := strings.TrimSpace(line)
		if s == "" && def != nil {
			return *def, nil
		}
		v, err := strconv.ParseFloat(s, 64)
		if err == nil {
			return v, nil
		}
		fmt.Fprintf(p.w, "数値として解釈できません: %q\n", s)
	}
}

func (p *Prompter) readLine() (string, error) {
	line, err := p.r.ReadString('\n')
	if err == io.EOF && line != "" {
		return line, nil
	}
	return line, err
}
