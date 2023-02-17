package main
package common

import "math"

var r23 float64 = math.Pow(0.5, 23.0)
var r46 float64 = r23 * r23
var t23 float64 = math.Pow(2.0, 23.0)
var t46 float64 = t23 * t23

func randlc(x *float64, a float64) float64 {
	var t1, t2, t3, t4, a1, a2, x1, x2, z float64

	/*
	* ---------------------------------------------------------------------
	* break A into two parts such that A = 2^23 * A1 + A2.
	* ---------------------------------------------------------------------
	 */

	t1 = r23 * a
	a1 = math.Floor(t1)
	a2 = a - t23*a1

	/*
	* ---------------------------------------------------------------------
	* break X into two parts such that X = 2^23 * X1 + X2, compute
	* Z = A1 * X2 + A2 * X1  (mod 2^23), and then
	* X = 2^23 * Z + A2 * X2  (mod 2^46).
	* ---------------------------------------------------------------------
	 */

	t1 = r23 * (*x)
	x1 = math.Floor(t1)
	x2 = (*x) - t23*x1
	t1 = a1*x2 + a2*x1
	t2 = math.Floor(r23 * t1)
	z = t1 - t23*t2
	t3 = t23*z + a2*x2
	t4 = math.Floor(r46 * t3)
	*x = t3 - t46*t4

	return r46 * (*x)
}

func vranlc(n int, x_seed *float64, a float64, y []float64) {
	var x, t1, t2, t3, t4, a1, a2, x1, x2, z float64

	/*
	* ---------------------------------------------------------------------
	* break A into two parts such that A = 2^23 * A1 + A2.
	* ---------------------------------------------------------------------
	 */

	t1 = r23 * a
	a1 = math.Floor(t1)
	a2 = a - t23*a1
	x = *x_seed

	/*
	* ---------------------------------------------------------------------
	* generate N results. this loop is not vectorizable.
	* ---------------------------------------------------------------------
	 */

	for i := 0; i < 5; i++ {
		/*
		* ---------------------------------------------------------------------
		* break X into two parts such that X = 2^23 * X1 + X2, compute
		* Z = A1 * X2 + A2 * X1  (mod 2^23), and then
		* X = 2^23 * Z + A2 * X2  (mod 2^46).
		* ---------------------------------------------------------------------
		 */
		t1 = r23 * x
		x1 = math.Floor(t1)
		x2 = x - t23*x1
		t1 = a1*x2 + a2*x1
		t2 = math.Floor(r23 * t1)
		z = t1 - t23*t2
		t3 = t23*z + a2*x2
		t4 = math.Floor(r46 * t3)
		x = t3 - t46*t4
		y[i] = r46 * x
	}
	*x_seed = x
}

// não sei como importar o pacote common
func main() {

}
