# Cletandup
## ✨ How To Use?
### 1. 채널에 등록하기
```
/cletanup addChannel
```
위 명령어를 입력하면 채널에 매일매일 **오전 10시쯤**에 오늘 날짜와 등록한 유저들에게 스탠드업 DM이 전송됩니다.

<br>


### 2. 유저 등록하기
```
/cletanup apply
```
유저 본인만 등록할 수 있으며, add된 채널에 해당 명령어를 입력하면 매일 **오전 10시쯤**에 apply한 유저 대상으로 DM을 보냅니다.

![image](https://user-images.githubusercontent.com/42836576/138896287-5a2f615d-aecd-45ce-86db-d038621fde7e.png) <br>
입력한 메시지는 해당 채널에 전송됩니다. (아직은 한꺼번에 답변을 ... 해야합니다 😭)

![image](https://user-images.githubusercontent.com/42836576/138896919-b04d96d0-193a-4e10-b105-d260fe538dd1.png) <br>
12시까지 입력하지 않으면 리마인드 DM을 전송합니다.


<br>

### 3. 채널 삭제하기
```
/cletanup deleteChannel
```


## 🌧 Build
```
git clone https://github.com/Lubycon/clelab-standup-mattermost.git
```

Build plugin:
```
make
```

This will produce a single plugin file (with support for multiple architectures) for upload to your Mattermost server:

```
dist/com.clelab.standup-0.1.0.tar.gz
```
