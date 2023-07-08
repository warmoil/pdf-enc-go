# pdf 암호화기입니다.

## 사용 방법 
in에는 변환할 pdf를 파일이름.pdf 이름으로 넣어야합니다
그리고 pw/pass.txt에는 파일이름,비밀번호로 를 적으면 됩니다.
그리고 프로그램을 실행하면 out폴더에 결과물이나옵니다

pass.txt 에는 이름,비밀번호가 있지만 in폴더에 해당 이름의 pdf가 없을경우에는
해당 이름,비밀번호는 무시하고 변환합니다.


build 명령어는 window 64기준으로

SET GOOS=windows& SET GOARCH=amd64& go build -o bin/app-windows-amd64.exe main.go

[그외빌드는_여기를_참고](https://dadev.tistory.com/entry/GO-Windows-macOS-%EB%B0%8F-Linux%EC%9A%A9-Go-%ED%94%84%EB%A1%9C%EA%B7%B8%EB%9E%A8%EC%9D%84-%EA%B5%90%EC%B0%A8-%EC%BB%B4%ED%8C%8C%EC%9D%BC%ED%95%98%EB%8A%94-%EB%B0%A9%EB%B2%95)