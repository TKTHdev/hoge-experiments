package main 


func DotNaive(a, b []float32) float32{
    sum := float32(0)
    for i := 0; i < len(a); i++ {
        sum += a[i] * b[i]
    }
    return sum
}

func DotUnroll4(a, b []float32) float32{
    sum := float32(0)
    for i:=0; i < len(a); i += 4 {
        s0 := a[i] * b[i]
        s1 := a[i+1] * b[i+1]
        s2 := a[i+2] * b[i+2]
        s3 := a[i+3] * b[i+3]
        sum += s0 + s1 + s2 + s3
    }
    return sum
}

func DotUnroll8(a, b []float32) float32{
    sum := float32(0)
    for i:=0; i < len(a); i += 8 {
        s0 := a[i] * b[i]
        s1 := a[i+1] * b[i+1]
        s2 := a[i+2] * b[i+2]
        s3 := a[i+3] * b[i+3]
        s4 := a[i+4] * b[i+4]
        s5 := a[i+5] * b[i+5]
        s6 := a[i+6] * b[i+6]
        s7 := a[i+7] * b[i+7]
        sum += s0 + s1 + s2 + s3 + s4 + s5 + s6 + s7
    }
    return sum
}

func DotUnroll4Optimized(a, b []float32) float32 {
    sum0 := float32(0)
    sum1 := float32(0)
    sum2 := float32(0)
    sum3 := float32(0)

    // 余りが出ないように、4の倍数までループ
    length := len(a) & ^3 
    
    for i := 0; i < length; i += 4 {
        // 【重要】BCEヒント: ここで範囲チェックを1回済ませる
        // これにより、以下の行での境界チェックが削除される
        _ = a[i+3]
        _ = b[i+3]

        // 【重要】変数を分けて依存関係を切る
        sum0 += a[i]   * b[i]
        sum1 += a[i+1] * b[i+1]
        sum2 += a[i+2] * b[i+2]
        sum3 += a[i+3] * b[i+3]
    }

    // 最後にまとめる
    total := sum0 + sum1 + sum2 + sum3

    // 端数（Tail）の処理
    for i := length; i < len(a); i++ {
        total += a[i] * b[i]
    }
    return total
}