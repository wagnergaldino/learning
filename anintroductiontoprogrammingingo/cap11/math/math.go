package math

func Average(xs []float64) float64 {
  total := float64(0)
  for _, x := range xs {
    total += x
  }
  return total / float64(len(xs))
}

func Min(xs []float64) float64 {
  min := float64(0)
  for i, x := range xs {
    if i == 0 || x < min {
    	min = x
    }
  }
  return min
}

func Max(xs []float64) float64 {
  max := float64(0)
  for i, x := range xs {
    if i == 0 || x > max {
    	max = x
    }
  }
  return max
}
