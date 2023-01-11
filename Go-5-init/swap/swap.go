package swap

//go约定大写开头的标识符才能被import用
func Swap(pValue1 *int, pValue2 *int) (swapResultMessage string) {
	temp := *pValue1
	*pValue1 = *pValue2
	*pValue2 = temp
	swapResultMessage = "swap success"
	return
}