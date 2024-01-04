package anyflag

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/spf13/pflag"
)

func setUpDSFlagSet(dsp *[]time.Duration) *pflag.FlagSet {
	f := pflag.NewFlagSet("test", pflag.ContinueOnError)
	f.VarP(NewSliceValue[time.Duration]([]time.Duration{}, dsp, time.ParseDuration), "ds", "", "Command separated list!")
	return f
}

func setUpDSFlagSetWithDefault(dsp *[]time.Duration) *pflag.FlagSet {
	f := pflag.NewFlagSet("test", pflag.ContinueOnError)
	f.VarP(NewSliceValue[time.Duration]([]time.Duration{0, 1}, dsp, time.ParseDuration), "ds", "", "Command separated list!")
	return f
}

func setUpDSFlagSetWithRedact(dsp *[]time.Duration) *pflag.FlagSet {
	r := func(d time.Duration) string {
		return d.Round(10 * time.Minute).String()
	}
	f := pflag.NewFlagSet("test", pflag.ContinueOnError)
	f.VarP(NewSliceValueWithRedact[time.Duration]([]time.Duration{}, dsp, time.ParseDuration, r), "ds", "", "Command separated list!")
	return f
}

func setUpSSFlagSet(ssp *[]string) *pflag.FlagSet {
	f := pflag.NewFlagSet("test", pflag.ContinueOnError)
	f.StringSliceVar(ssp, "ss", []string{}, "Command separated list!")
	return f
}

func TestEmptyDS(t *testing.T) {
	var ds []time.Duration
	f := setUpDSFlagSet(&ds)
	err := f.Parse([]string{})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}

	if len(ds) != 0 {
		t.Fatalf("got ds %v with len=%d but expected length=0", ds, len(ds))
	}
}

func TestDS(t *testing.T) {
	var ds []time.Duration
	f := setUpDSFlagSet(&ds)

	vals := []string{"1ns", "2ms", "3m", "4h"}
	arg := fmt.Sprintf("--ds=%s", strings.Join(vals, ","))
	err := f.Parse([]string{arg})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}
	for i, v := range ds {
		d, err := time.ParseDuration(vals[i])
		if err != nil {
			t.Fatalf("got error: %v", err)
		}
		if d != v {
			t.Fatalf("expected ds[%d] to be %s but got: %d", i, vals[i], v)
		}
	}
}

func TestDSDefault(t *testing.T) {
	var ds []time.Duration
	f := setUpDSFlagSetWithDefault(&ds)

	vals := []string{"0s", "1ns"}

	err := f.Parse([]string{})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}
	for i, v := range ds {
		d, err := time.ParseDuration(vals[i])
		if err != nil {
			t.Fatalf("got error: %v", err)
		}
		if d != v {
			t.Fatalf("expected ds[%d] to be %d but got: %d", i, d, v)
		}
	}
}

func TestDSWithDefault(t *testing.T) {
	var ds []time.Duration
	f := setUpDSFlagSetWithDefault(&ds)

	vals := []string{"1ns", "2ns"}
	arg := fmt.Sprintf("--ds=%s", strings.Join(vals, ","))
	err := f.Parse([]string{arg})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}
	for i, v := range ds {
		d, err := time.ParseDuration(vals[i])
		if err != nil {
			t.Fatalf("got error: %v", err)
		}
		if d != v {
			t.Fatalf("expected ds[%d] to be %d but got: %d", i, d, v)
		}
	}
}

func TestDSWithRedact(t *testing.T) {
	var ds []time.Duration
	f := setUpDSFlagSetWithRedact(&ds)

	vals := []string{"1ns", "2ms", "3m", "4h"}
	arg := fmt.Sprintf("--ds=%s", strings.Join(vals, ","))
	err := f.Parse([]string{arg})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}

	v := f.Lookup("ds").Value
	if v.String() != "[0s,0s,0s,4h0m0s]" {
		t.Fatal("expected redacted; got", v.String())
	}
	if s := v.(*SliceValue[time.Duration]); strings.Join(s.GetSlice(), ",") != "0s,0s,0s,4h0m0s" {
		t.Fatal("expected redacted; got", s.GetSlice())
	}
}

func TestDSAsSliceValue(t *testing.T) {
	var ds []time.Duration
	f := setUpDSFlagSet(&ds)

	in := []string{"1ns", "2ns"}
	argfmt := "--ds=%s"
	arg1 := fmt.Sprintf(argfmt, in[0])
	arg2 := fmt.Sprintf(argfmt, in[1])
	err := f.Parse([]string{arg1, arg2})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}

	f.VisitAll(func(f *pflag.Flag) {
		if val, ok := f.Value.(*SliceValue[time.Duration]); ok {
			_ = val.Replace([]string{"3ns"})
		}
	})
	if len(ds) != 1 || ds[0] != time.Duration(3) {
		t.Fatalf("Expected ss to be overwritten with '3ns', but got: %v", ds)
	}
}

func TestDSCalledTwice(t *testing.T) {
	var ds []time.Duration
	f := setUpDSFlagSet(&ds)

	in := []string{"1ns,2ns", "3ns"}
	expected := []time.Duration{1, 2, 3}
	argfmt := "--ds=%s"
	arg1 := fmt.Sprintf(argfmt, in[0])
	arg2 := fmt.Sprintf(argfmt, in[1])
	err := f.Parse([]string{arg1, arg2})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}
	for i, v := range ds {
		if expected[i] != v {
			t.Fatalf("expected ds[%d] to be %d but got: %d", i, expected[i], v)
		}
	}
}

func TestWithCommas(t *testing.T) {
	var ss []string

	f := setUpSSFlagSet(&ss)
	in := []string{`"foo=1,2"`, `"bar=3"`}
	expected := []string{"foo=1,2", "bar=3"}
	argfmt := "--ss=%s"
	arg1 := fmt.Sprintf(argfmt, in[0])
	arg2 := fmt.Sprintf(argfmt, in[1])
	err := f.Parse([]string{arg1, arg2})
	if err != nil {
		t.Fatal("expected no error; got", err)

	}
	for i, v := range ss {
		if expected[i] != v {
			t.Fatalf("expected ss[%d] to be %s but got: %s", i, expected[i], v)

		}

	}
}
