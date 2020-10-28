package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Signal 串口接收的单条数据
type Signal struct {
	Seq       string    // 串口的一行文本记录
	ColumnCnt int       // 多少个字段
	Columns   []float64 // 字段数据
	Model     string    // 型号 15 or 16 > co2，10 or 11 > no_co2
}

// ParserSignal 解析串口数据
func ParserSignal(input string) (*Signal, error) {
	columnsStr := strings.Split(input, " ")
	if len(columnsStr) < 0 {
		return nil, fmt.Errorf("忽略，不是data: %s", input)
	}
	if len(columnsStr) < 16 {
		return nil, fmt.Errorf("忽略，未定义的数据长度＜ %s", input)
	}
	if len(columnsStr) > 17 {
		return nil, fmt.Errorf("忽略，未定义的数据长度 > %s", input)
	}
	var columns []float64
	for _, f := range columnsStr {
		fi, err := strconv.ParseFloat(f, 64)
		columns = append(columns, fi)
		if err != nil {
			return nil, fmt.Errorf("忽略，数据字段类型不对 > %s", input)
		}
	}

	// 16位字段数据不带CO2版本
	if len(columns) == 16 {
		return &Signal{
			Seq:       input,
			ColumnCnt: len(columns),
			Columns:   columns,
			Model:     "no_co2",
		}, nil
	}

	// 16位字段数据不带CO2版本
	if len(columns) == 17 {
		return &Signal{
			Seq:       input,
			ColumnCnt: len(columns),
			Columns:   columns,
			Model:     "co2",
		}, nil
	}
	return nil, errors.New("错误，未定义的数据解析")
}
