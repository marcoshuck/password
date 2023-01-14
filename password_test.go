package password

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func BenchmarkValidatePassword(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Validate("%+f9mi4#S,o02exo^MC=zkcrkOpf12VMUAXjP+8LBSICO.{vJO")
	}
}

func TestPasswordShouldBeAtLeast8CharsLong(t *testing.T) {
	err := Validate("a1b2")
	require.Error(t, err)
	assert.Equal(t, ErrLength8, err)
}

func TestPasswordShouldHaveAtLeast2NumericChars(t *testing.T) {
	err := Validate("%+fmi#S,oexo")
	require.Error(t, err)
	assert.Equal(t, ErrMissingNumeric, err)
}

func TestPasswordShouldHaveAtLeast4AlphabeticChars(t *testing.T) {
	err := Validate(";^82&!&[877=,,.+@)2=+8+5$[!;(8=3~#;&0-3~{105++~]~^")
	require.Error(t, err)
	assert.Equal(t, ErrMissingAlphabetic, err)
}

func TestPasswordShouldHaveAtLeast1SpecialChar(t *testing.T) {
	err := Validate("cSu8hMY5Y4iBWs7RzPiogYKZSwbZcKP0lyBsV2bXxr1Gy9J2Tk")
	require.Error(t, err)
	assert.Equal(t, ErrMissingSpecial, err)
}

func TestPasswordShouldHaveAtLeast1UppercaseChar(t *testing.T) {
	err := Validate("~6}pkhto=5jvtfss$izz!e2p3mw@}&hjz.v_eh321@&%jj8s4;")
	require.Error(t, err)
	assert.Equal(t, ErrMissingUppercase, err)
}

func TestPasswordShouldNotHaveQwerty(t *testing.T) {
	err := Validate("qwerty%+f9mi4#S,o02exo^MC=zkcrkOpf12VMUAXjP+8LBSICO.{vJO")
	require.Error(t, err)
	assert.Equal(t, ErrInvalidCharCombination, err)
}

func TestPasswordShouldNotHaveAsdf(t *testing.T) {
	err := Validate("asdf%+f9mi4#S,o02exo^MC=zkcrkOpf12VMUAXjP+8LBSICO.{vJO")
	require.Error(t, err)
	assert.Equal(t, ErrInvalidCharCombination, err)
}

func TestPasswordShouldNotHave1234(t *testing.T) {
	err := Validate("1234%+f9mi4#S,o02exo^MC=zkcrkOpf12VMUAXjP+8LBSICO.{vJO")
	require.Error(t, err)
	assert.Equal(t, ErrConsecutive, err)
}

func TestPasswordShouldNotHave98765(t *testing.T) {
	err := Validate("98765%+f9mi4#S,o02exo^MC=zkcrkOpf12VMUAXjP+8LBSICO.{vJO")
	require.Error(t, err)
	assert.Equal(t, ErrConsecutive, err)
}

func TestPasswordShouldNotHaveEqualAdjacent(t *testing.T) {
	err := Validate("%+f9mi4#S,o02exoooo^MC=zkcrkOpf12VMUAXjP+8LBSICO.{vJO")
	require.Error(t, err)
	assert.Equal(t, ErrEqualAdjacent, err)
}
