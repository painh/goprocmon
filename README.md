config.json.sample 을 복사하여
config.json 으로 만들어주세요

{
  "port": "9991", //실행 포트
  "processes": {
    "firefox": { "proc": "firefox.exe", "ret": "cnt" },
    "code": { "proc": "Code", "ret": "on" }
  }
}

proc 는 검사할 프로세스 이름입니다. 앞부분이 부분일치하면 통과합니다.
ret는
    cnt : 갯수
    on : on/off
로 표시합니다.


실행은

리눅스 
chmod +x main
./main

윈도우
main.exe

로 해주세요.