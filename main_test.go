package main

import "testing"

func BenchmarkGetFromMap(b *testing.B) {
	str := `{
    "id": 1,
    "name": "John Doe",
    "age": 29,
    "gender": "male",
    "email": "johndoe@example.com",
    "phone": "123-456-7890",
    "address": "123 Main St, Anytown, USA",
    "city": "Anytown",
    "state": "CA",
    "zip_code": "12345",
    "country": "USA",
    "username": "johndoe123",
    "password": "s3cur3P@ssw0rd",
    "ssn": "123-45-6789",
    "birthdate": "1994-01-01",
    "company": "TechCorp Inc.",
    "job_title": "Software Engineer",
    "department": "Engineering",
    "salary": 85000,
    "hire_date": "2020-06-15",
    "is_full_time": true,
    "is_remote": false,
    "last_login": "2024-09-04T10:00:00Z",
    "preferred_language": "English",
    "timezone": "PST",
    "manager_id": 5,
    "emergency_contact_name": "Jane Doe",
    "emergency_contact_phone": "987-654-3210",
    "emergency_contact_relation": "Spouse",
    "profile_picture_url": "https://example.com/images/johndoe.jpg",
    "bio": "Software engineer with 5+ years of experience in full-stack development.",
    "website": "https://johndoe.dev",
    "linkedin_url": "https://linkedin.com/in/johndoe",
    "github_url": "https://github.com/johndoe",
    "twitter_handle": "@johndoe",
    "skills": ["Go", "JavaScript", "React", "Node.js"],
    "certifications": ["AWS Certified Developer", "Certified Kubernetes Administrator"],
    "education": "B.Sc. in Computer Science",
    "hobbies": ["hiking", "photography", "gaming"],
    "favorite_quote": "Code is like humor. When you have to explain it, it’s bad."
}
`

	for i := 0; i < b.N; i++ {
		GetFromMap(str, "id", "name", "age", "gender", "email", "phone", "address", "city", "state",
			"zip_code", "country", "username", "password", "password", "ssn", "birthdate", "company", "job_title",
			"department", "salary", "hire_date", "is_full_time", "is_remote", "last_login", "preferred_language",
			"timezone", "manager_id", "emergency_contact_name", "emergency_contact_phone", "emergency_contact_relation",
			"profile_picture_url", "bio", "website", "linkedin_url", "github_url", "twitter_handle", "skills",
			"certifications", "education", "hobbies", "favorite_quote")
	}
}

//func BenchmarkGetStr(b *testing.B) {
//	str := `{
//  "name": "John Doe",
//  "age": 30,
//  "email": "johndoe@example.com",
//  "city": "New York",
//  "isEmployed": true
//}`
//
//	for i := 0; i < b.N; i++ {
//		GetStr(str, "name", "age", "email", "city", "isEmployed")
//	}
//}

func BenchmarkGetOne(b *testing.B) {
	str := `{
    "id": 1,
    "name": "John Doe",
    "age": 29,
    "gender": "male",
    "email": "johndoe@example.com",
    "phone": "123-456-7890",
    "address": "123 Main St, Anytown, USA",
    "city": "Anytown",
    "state": "CA",
    "zip_code": "12345",
    "country": "USA",
    "username": "johndoe123",
    "password": "s3cur3P@ssw0rd",
    "ssn": "123-45-6789",
    "birthdate": "1994-01-01",
    "company": "TechCorp Inc.",
    "job_title": "Software Engineer",
    "department": "Engineering",
    "salary": 85000,
    "hire_date": "2020-06-15",
    "is_full_time": true,
    "is_remote": false,
    "last_login": "2024-09-04T10:00:00Z",
    "preferred_language": "English",
    "timezone": "PST",
    "manager_id": 5,
    "emergency_contact_name": "Jane Doe",
    "emergency_contact_phone": "987-654-3210",
    "emergency_contact_relation": "Spouse",
    "profile_picture_url": "https://example.com/images/johndoe.jpg",
    "bio": "Software engineer with 5+ years of experience in full-stack development.",
    "website": "https://johndoe.dev",
    "linkedin_url": "https://linkedin.com/in/johndoe",
    "github_url": "https://github.com/johndoe",
    "twitter_handle": "@johndoe",
    "skills": ["Go", "JavaScript", "React", "Node.js"],
    "certifications": ["AWS Certified Developer", "Certified Kubernetes Administrator"],
    "education": "B.Sc. in Computer Science",
    "hobbies": ["hiking", "photography", "gaming"],
    "favorite_quote": "Code is like humor. When you have to explain it, it’s bad."
}
`

	for i := 0; i < b.N; i++ {
		GetOne(str, "name")
	}
}
