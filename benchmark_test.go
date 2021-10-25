package main

import (
	"testing"
)

func BenchmarkGenXid(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		genXid()
	}
}

func BenchmarkGenKsuid(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		genKsuid()
	}
}

func BenchmarkGenBetterGUID(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		genBetterGUID()
	}
}

func BenchmarkGenUlidD(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		genUlid()
	}
}

func BenchmarkGenSonyflake(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		genSonyflake()
	}
}

func BenchmarkGenSid(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		genSid()
	}
}

func BenchmarkGenShortUUID(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		genShortUUID()
	}
}

func BenchmarkGenUUIDv4(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		genUUIDv4()
	}
}

func BenchmarkGenUUID1v4(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		genUUID1v4()
	}
}

func BenchmarkGenUUID(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		genUUID()
	}
}

func BenchmarkGenCUID(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		genCUID()
	}
}

func BenchmarkGenNanoID(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		genNanoID()
	}
}
