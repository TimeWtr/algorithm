package main

import "fmt"

// 分治法求最大方块尺寸（核心算法）
func findLargestSquare(length, width int) int {
	// 确保长度大于宽度，方便计算
	if width > length {
		length, width = width, length
	}

	// 基线条件：当宽度能整除长度时
	if length%width == 0 {
		return width
	}

	// 递归分解：从土地中切出最大正方形后处理剩余部分
	return findLargestSquare(width, length%width)
}

// 计算划分方案详情（方块尺寸、数量、小矩形信息）
func calculateDivisionPlan(length, width int) (squareSize, count int, remainingWidth, remainingLength int) {
	// 获取最大方块尺寸
	squareSize = findLargestSquare(length, width)

	// 计算完整方块数量
	count = (length / squareSize) * (width / squareSize)

	// 计算剩余土地尺寸
	remainingLength = length % squareSize
	remainingWidth = width % squareSize

	return
}

// 可视化显示划分方案
func visualizeDivision(length, width, squareSize int) {
	fmt.Printf("\n土地尺寸: %d × %d\n", length, width)
	fmt.Printf("最大方块尺寸: %d × %d\n", squareSize, squareSize)

	// 计算网格行列数
	rows := width / squareSize
	cols := length / squareSize
	remainingWidth := width % squareSize
	remainingLength := length % squareSize

	// 打印网格
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Print("■ ")
		}
		// 打印右侧剩余部分
		if remainingLength > 0 {
			fmt.Printf("▦ (剩余部分 %d×%d)", remainingLength, squareSize)
		}
		fmt.Println()
	}

	// 打印底部剩余部分
	if remainingWidth > 0 {
		for j := 0; j < cols; j++ {
			fmt.Print("▦ ")
		}
		if remainingLength > 0 {
			fmt.Printf("▦ (剩余部分 %d×%d)", remainingLength, remainingWidth)
		}
		fmt.Println()
	}
}

func main() {
	length, width := 1680, 640

	// 计算最大方块尺寸
	squareSize := findLargestSquare(length, width)
	fmt.Printf("===== 土地划分结果: %d×%d → 方块尺寸 %d×%d =====\n",
		length, width, squareSize, squareSize)

	// 获取划分详情
	squareSize, count, remW, remL := calculateDivisionPlan(length, width)
	fmt.Printf("完整方块数量: %d\n", count)
	fmt.Printf("剩余土地尺寸: %d×%d\n", remL, remW)

	// 可视化显示
	visualizeDivision(length, width, squareSize)

	// 更多测试用例
	testCases := [][2]int{
		{640, 400},  // 640×400 -> 80×80
		{400, 240},  // 400×240 -> 80×80
		{240, 160},  // 240×160 -> 80×80
		{160, 80},   // 160×80  -> 80×80
		{100, 50},   // 100×50  -> 50×50
		{75, 45},    // 75×45   -> 15×15
		{1280, 720}, // 1280×720 -> 80×80
		{100, 30},   // 100×30  -> 10×10
	}

	fmt.Println("\n更多测试用例验证:")
	for _, tc := range testCases {
		size := findLargestSquare(tc[0], tc[1])
		fmt.Printf("土地 %d×%d → 方块 %d×%d\n", tc[0], tc[1], size, size)
	}
}
