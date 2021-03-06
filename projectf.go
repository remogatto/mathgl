package mathgl

import (
	"errors"
	"math"
)

func Ortho(left, right, bottom, top, near, far float32) Mat4f {
	rml, tmb, fmn := (right - left), (top - bottom), (far - near)

	return Mat4f{float32(2. / rml), 0, 0, 0, 0, float32(2. / tmb), 0, 0, 0, 0, float32(-2. / fmn), 0, float32(-(right + left) / rml), float32(-(top + bottom) / tmb), float32(-(far + near) / fmn), 1}
}

// Equivalent to Ortho with the near and far planes being -1 and 1, respectively
func Ortho2D(left, right, top, bottom float32) Mat4f {
	return Ortho(left, right, top, bottom, -1, 1)
}

func Perspective(fovy, aspect, near, far float32) Mat4f {
	fovy = (fovy * math.Pi) / 180.0 // convert from degrees to radians
	nmf, f := near-far, float32(1./math.Tan(float64(fovy)/2.0))

	return Mat4f{float32(f / aspect), 0, 0, 0, 0, float32(f), 0, 0, 0, 0, float32((near + far) / nmf), -1, 0, 0, float32((2. * far * near) / nmf), 0}
}

func Frustum(left, right, bottom, top, near, far float32) Mat4f {
	rml, tmb, fmn := (right - left), (top - bottom), (far - near)
	A, B, C, D := (right+left)/rml, (top+bottom)/tmb, -(far+near)/fmn, (2*far*near)/fmn

	return Mat4f{float32((2. * near) / rml), 0, 0, 0, 0, float32((2. * near) / tmb), 0, 0, float32(A), float32(B), float32(C), -1, 0, 0, float32(D), 0}
}

func LookAt(eyeX, eyeY, eyeZ, centerX, centerY, centerZ, upX, upY, upZ float32) Mat4f {
	F := Vec3f{
		float32(centerX - eyeX),
		float32(centerY - eyeY),
		float32(centerZ - eyeZ)}

	f := F.Normalize()

	Up := Vec3f{
		float32(upX),
		float32(upY),
		float32(upZ)}

	Upp := Up.Normalize()

	s := f.Cross(Upp)
	u := s.Cross(f)

	M := Mat4f{s[0], u[0], -f[0], 0, s[1], u[1], -f[1], 0, s[2], u[2], -f[2], 0, 0, 0, 0, 1}

	return M.Mul4(Translate3D(-eyeX, -eyeY, -eyeZ))
}

func LookAtV(eye, center, up Vec3f) Mat4f {
	F := center.Sub(eye)

	f := F.Normalize()

	Upp := up.Normalize()

	s := f.Cross(Upp)
	u := s.Cross(f)

	M := Mat4f{s[0], u[0], -f[0], 0, s[1], u[1], -f[1], 0, s[2], u[2], -f[2], 0, 0, 0, 0, 1}

	return M.Mul4(Translate3D(float32(-eye[0]), float32(-eye[1]), float32(-eye[2])))
}

// Transform a set of coordinates from object space (in obj) to window coordinates (with depth)
//
// Window coordinates are continuous, not discrete (well, as continuous as an IEEE Floating Point can be), so you won't get exact pixel locations
// without rounding or similar
func Projectf(obj Vec3f, modelview, projection Mat4f, initialX, initialY, width, height int) (win Vec3f) {
	obj4 := Vec4f{obj[0], obj[1], obj[2], 1.0}

	vpp := projection.Mul4(modelview).Mul4x1(obj4)
	win[0] = float32(initialX) + (float32(width)*(vpp[0]+1))/2
	win[1] = float32(initialY) + (float32(height)*(vpp[1]+1))/2
	win[2] = (vpp[2] + 1) / 2

	return win
}

// Transform a set of window coordinates to object space. If your MVP (projection.Mul(modelview) matrix is not invertible, this will return an error
//
// Note that the projection may not be perfect if you use strict pixel locations rather than the exact values given by Projectf.
// (It's still unlikely to be perfect due to precision errors, but it will be closer)
func UnProjectf(win Vec3f, modelview, projection Mat4f, initialX, initialY, width, height int) (obj Vec3f, err error) {
	inv := projection.Mul4(modelview).Inv()
	blank := Mat4f{}
	if inv == blank {
		return Vec3f{}, errors.New("Could not find matrix inverse (projection times modelview is probably non-singular)")
	}

	obj[0] = (2 * (win[0] - float32(initialX)) / float32(width)) - 1
	obj[1] = (2 * (win[1] - float32(initialY)) / float32(height)) - 1
	obj[2] = 2*win[2] - 1

	return obj, nil
}
