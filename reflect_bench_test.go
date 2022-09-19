package main

import (
	"reflect"
	"testing"
)

type ShortStruct struct {
	A string
}

type ShortUpdate struct {
	A *string
}

type Struct struct {
	A, B, C, D, E string
}

type Update struct {
	A, B, C, D, E *string
}

type LongStruct struct {
	SomeFieldName, AnotherFieldName, ALongFieldName, YetAnotherFieldName, A, B, C, D, E, F, TheFieldNameIWant string
}

type LongUpdate struct {
	SomeFieldName, AnotherFieldName, ALongFieldName, YetAnotherFieldName, A, B, C, D, E, F, TheFieldNameIWant *string
}

func BenchmarkNativeSetShort(b *testing.B) {
	v := &ShortStruct{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.A = "SomeString"
	}
}

func BenchmarkNativeSet(b *testing.B) {
	v := &Struct{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.C = "SomeString"
	}
}

func BenchmarkNativeSetLong(b *testing.B) {
	v := &LongStruct{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.TheFieldNameIWant = "SomeString"
	}
}

func BenchmarkReflectSetByNameShort(b *testing.B) {
	ReflectSetByName(b, &ShortStruct{}, "A")
}

func BenchmarkReflectSetByName(b *testing.B) {
	ReflectSetByName(b, &Struct{}, "C")
}

func BenchmarkReflectSetByNameLong(b *testing.B) {
	ReflectSetByName(b, &LongStruct{}, "TheFieldNameIWant")
}

func ReflectSetByName(b *testing.B, obj any, fieldName string) {
	for i := 0; i < b.N; i++ {
		v := reflect.ValueOf(obj).Elem()
		field := v.FieldByName(fieldName)
		field.SetString("SomeString")
	}
}

func BenchmarkReflectSetByIndexShort(b *testing.B) {
	ReflectSetByIndex(b, &ShortStruct{}, 0)
}

func BenchmarkReflectSetByIndex(b *testing.B) {
	ReflectSetByIndex(b, &Struct{}, 2)
}

func BenchmarkReflectSetByIndexLong(b *testing.B) {
	ReflectSetByIndex(b, &LongStruct{}, 10)
}

func ReflectSetByIndex(b *testing.B, obj any, fieldIndex int) {
	for i := 0; i < b.N; i++ {
		v := reflect.ValueOf(obj).Elem()
		field := v.Field(fieldIndex)
		field.SetString("SomeString")
	}
}

func BenchmarkReflectSetByNameCachedShort(b *testing.B) {
	ReflectSetByNameCached(b, &ShortStruct{}, "A")
}

func BenchmarkReflectSetByNameCached(b *testing.B) {
	ReflectSetByNameCached(b, &Struct{}, "C")
}

func BenchmarkReflectSetByNameCachedLong(b *testing.B) {
	ReflectSetByNameCached(b, &LongStruct{}, "TheFieldNameIWant")
}

func ReflectSetByNameCached(b *testing.B, obj any, fieldName string) {
	cache := make(map[string]int)

	v := reflect.ValueOf(obj).Elem()
	t := v.Type()
	numField := t.NumField()
	for i := 0; i < numField; i++ {
		cache[t.Field(i).Name] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		field := v.Field(cache[fieldName])
		field.SetString("SomeString")
	}
}

func BenchmarkNativeUpdateShort(b *testing.B) {
	v := &ShortStruct{}
	someString := "SomeString"
	u := &ShortUpdate{A: &someString}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		setINN(u.A, &v.A)
	}
}

func BenchmarkNativeUpdate(b *testing.B) {
	v := &Struct{}
	someString := "SomeString"
	u := &Update{
		A: &someString,
		B: &someString,
		C: &someString,
		D: &someString,
		E: &someString,
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		setINN(u.A, &v.A)
		setINN(u.B, &v.B)
		setINN(u.C, &v.C)
		setINN(u.D, &v.D)
		setINN(u.E, &v.E)
	}
}

func BenchmarkNativeUpdateLong(b *testing.B) {
	v := &LongStruct{}
	someString := "SomeString"
	u := &LongUpdate{
		SomeFieldName:       &someString,
		AnotherFieldName:    &someString,
		ALongFieldName:      &someString,
		YetAnotherFieldName: &someString,
		A:                   &someString,
		B:                   &someString,
		C:                   &someString,
		D:                   &someString,
		E:                   &someString,
		F:                   &someString,
		TheFieldNameIWant:   &someString,
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		setINN(u.SomeFieldName, &v.SomeFieldName)
		setINN(u.AnotherFieldName, &v.AnotherFieldName)
		setINN(u.ALongFieldName, &v.ALongFieldName)
		setINN(u.YetAnotherFieldName, &v.YetAnotherFieldName)
		setINN(u.A, &v.A)
		setINN(u.B, &v.B)
		setINN(u.C, &v.C)
		setINN(u.D, &v.D)
		setINN(u.E, &v.E)
		setINN(u.F, &v.F)
		setINN(u.SomeFieldName, &v.SomeFieldName)
	}
}

func setINN(from, to *string) {
	if from != nil {
		*to = *from
	}
}

func BenchmarkReflectUpdateByNameShort(b *testing.B) {
	v := &ShortStruct{}
	someString := "SomeString"
	u := &ShortUpdate{A: &someString}
	b.ResetTimer()
	ReflectUpdateByName(b, v, u)
}

func BenchmarkReflectUpdateByName(b *testing.B) {
	v := &Struct{}
	someString := "SomeString"
	u := &Update{
		A: &someString,
		B: &someString,
		C: &someString,
		D: &someString,
		E: &someString,
	}
	b.ResetTimer()
	ReflectUpdateByName(b, v, u)
}

func BenchmarkReflectUpdateByNameLong(b *testing.B) {
	v := &LongStruct{}
	someString := "SomeString"
	u := &LongUpdate{
		SomeFieldName:       &someString,
		AnotherFieldName:    &someString,
		ALongFieldName:      &someString,
		YetAnotherFieldName: &someString,
		A:                   &someString,
		B:                   &someString,
		C:                   &someString,
		D:                   &someString,
		E:                   &someString,
		F:                   &someString,
		TheFieldNameIWant:   &someString,
	}
	b.ResetTimer()
	ReflectUpdateByName(b, v, u)
}

func ReflectUpdateByName(b *testing.B, obj any, update any) {
	v := reflect.ValueOf(obj).Elem()
	t := v.Type()
	updateV := reflect.ValueOf(update).Elem()
	updateT := updateV.Type()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		numField := updateT.NumField()
		for j := 0; j < numField; j++ {
			updateFieldName := updateT.Field(j).Name
			field, ok := t.FieldByName(updateFieldName)
			if !ok {
				continue
			}
			updateField := updateV.Field(j)
			if updateField.IsNil() {
				continue
			}

			v.FieldByIndex(field.Index).SetString(updateField.Elem().String())
		}
	}
}

func BenchmarkReflectUpdateByIndexShort(b *testing.B) {
	v := &ShortStruct{}
	someString := "SomeString"
	u := &ShortUpdate{A: &someString}
	b.ResetTimer()
	ReflectUpdateByIndex(b, v, u)
}

func BenchmarkReflectUpdateByIndex(b *testing.B) {
	v := &Struct{}
	someString := "SomeString"
	u := &Update{
		A: &someString,
		B: &someString,
		C: &someString,
		D: &someString,
		E: &someString,
	}
	b.ResetTimer()
	ReflectUpdateByIndex(b, v, u)
}

func BenchmarkReflectUpdateByIndexLong(b *testing.B) {
	v := &LongStruct{}
	someString := "SomeString"
	u := &LongUpdate{
		SomeFieldName:       &someString,
		AnotherFieldName:    &someString,
		ALongFieldName:      &someString,
		YetAnotherFieldName: &someString,
		A:                   &someString,
		B:                   &someString,
		C:                   &someString,
		D:                   &someString,
		E:                   &someString,
		F:                   &someString,
		TheFieldNameIWant:   &someString,
	}
	b.ResetTimer()
	ReflectUpdateByIndex(b, v, u)
}

func ReflectUpdateByIndex(b *testing.B, obj any, update any) {
	v := reflect.ValueOf(obj).Elem()
	updateV := reflect.ValueOf(update).Elem()
	updateT := updateV.Type()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		numField := updateT.NumField()
		for j := 0; j < numField; j++ {
			updateField := updateV.Field(j)
			if updateField.IsNil() {
				continue
			}

			v.Field(j).SetString(updateField.Elem().String())
		}
	}
}

func BenchmarkReflectUpdateByNameCachedShort(b *testing.B) {
	v := &ShortStruct{}
	someString := "SomeString"
	u := &ShortUpdate{A: &someString}
	b.ResetTimer()
	ReflectUpdateByNameCached(b, v, u)
}

func BenchmarkReflectUpdateByNameCached(b *testing.B) {
	v := &Struct{}
	someString := "SomeString"
	u := &Update{
		A: &someString,
		B: &someString,
		C: &someString,
		D: &someString,
		E: &someString,
	}
	b.ResetTimer()
	ReflectUpdateByNameCached(b, v, u)
}

func BenchmarkReflectUpdateByNameCachedLong(b *testing.B) {
	v := &LongStruct{}
	someString := "SomeString"
	u := &LongUpdate{
		SomeFieldName:       &someString,
		AnotherFieldName:    &someString,
		ALongFieldName:      &someString,
		YetAnotherFieldName: &someString,
		A:                   &someString,
		B:                   &someString,
		C:                   &someString,
		D:                   &someString,
		E:                   &someString,
		F:                   &someString,
		TheFieldNameIWant:   &someString,
	}
	b.ResetTimer()
	ReflectUpdateByNameCached(b, v, u)
}

func ReflectUpdateByNameCached(b *testing.B, obj any, update any) {
	type Pair struct{ vID, uID int }
	v := reflect.ValueOf(obj).Elem()
	t := v.Type()
	updateV := reflect.ValueOf(update).Elem()
	updateT := updateV.Type()

	numField := updateT.NumField()
	cache := make([]Pair, 0, numField)
	for i := 0; i < numField; i++ {
		fieldName := updateT.Field(i).Name
		field, ok := t.FieldByName(fieldName)
		if !ok {
			continue
		}
		cache = append(cache, Pair{
			vID: field.Index[0],
			uID: i,
		})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, pair := range cache {
			updateField := updateV.Field(pair.uID)
			if updateField.IsNil() {
				continue
			}

			v.Field(pair.vID).SetString(updateField.Elem().String())
		}
	}
}
