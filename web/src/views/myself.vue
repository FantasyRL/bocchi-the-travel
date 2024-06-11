<script setup>
import axios from "axios";
import Cookies from "js-cookie";
</script>

<script>
export default {
  data() {
    return {
      userid: "",
      setname: "设置",
      setcount: 3,
      setshow: 0,
      id: 1,
      name: "未登录",
      avatar: "//xiey.work/ava.png",
      mail: "",
      signature: "这个人很懒，什么都没写。",
      access_token: "",
      refresh_token: "",
      avatarFile: null,
      url: "",
      showavatarpage: 0,
      previewImage: null,
      file: null,
      isEditing: false
    };
  },
  methods: {
    toifollow() {
      this.$router.push(`/ifollow/${this.id}`);
    },
    tofollowme() {
      this.$router.push(`/followme/${this.id}`);
    },
    savesignature() {
      this.isEditing = false;
      axios
        .post(
          "https://api.xiey.work/bocchi/user/signature",
          {
            signature: this.signature
          },
          {
            headers: {
              "access-token": this.access_token
            }
          }
        )
        .then((res) => {
          console.log(res);
        })
        .catch((err) => {
          console.error(err);
        });
    },
    editsignature() {
      this.isEditing = true;
      const input = document.querySelector(".input");
      input.focus();
    },

    modavatarpage() {
      this.showavatarpage = !this.showavatarpage;
      console.log(this.showavatarpage);
    },
    onFileChange(e) {
      const file = e.target.files[0];
      this.file = file; // 保存文件对象，以便稍后上传。
      if (file) {
        const reader = new FileReader();
        reader.onload = (e) => {
          this.previewImage = e.target.result;
        };
        reader.readAsDataURL(file);
      }
    },
    putavatar(previewImage) {
      // 如果没有选择文件，则改变warm的值
      if (!previewImage) {
        this.warm = 1;
        return;
      }
      if (!previewImage.type.startsWith("image/")) {
        alert("请选择图片文件");
        return;
      }
      if (previewImage.size > 4 * 1024 * 1024) {
        alert("图片大小不能超过4MB");
        return;
      }

      var currentTime = new Date();
      var seconds = currentTime.getSeconds();

      console.log(seconds);
      axios
        .put(
          "https://api.xiey.work/bocchi/user/avatar/upload",
          {
            avatar_file: previewImage
          },
          {
            headers: {
              "Content-Type": "multipart/form-data", // 设置请求头，告诉服务器这是一个multipart/form-data请求，因为我们要上传文件。
              "access-token": this.access_token, // 假设返回的数据中有一个名为access_token的字段，表示访问令牌。
              Accept: "*/*" // 设置接收任意类型的响应。
            }
          }
        )
        .then((res) => {
          console.log(res);
          this.avatar = res.data.user.avatar + "?" + seconds;
          this.modavatarpage();
          this.warm = 1;
        })
        .catch((err) => {
          console.error(err);
        });
    },
    settings() {
      this.setshow = !this.setshow;
      this.setcount++; // 假设返回的数据中有一个名为name的字段，表示用户名。
    },
    refreshtoken() {
      axios
        .get("https://api.xiey.work/bocchi/access_token/get", {
          headers: {
            "refresh-token": this.refresh_token,
            Accept: "*/*"
          }
        })
        .then((res) => {
          console.log(res);
          console.log(res.data.access_token);
          this.access_token = res.data.access_token;
          Cookies.remove("access_token");
          this.cookiesSet = Cookies.set("access_token", res.data.access_token, {
            expires: 1
          });
        })
        .catch((err) => {
          console.error(err);
        });
    },

    logout() {
      Cookies.remove("id");
      Cookies.remove("access_token");
      Cookies.remove("refresh_token");
      window.location.href = "/";
    },
    init() {
      axios
        .get("https://api.xiey.work/bocchi/user/info?user_id=" + this.id, {
          headers: {
            "access-token": this.access_token
          }
        })
        .then((res) => {
          console.log(res);
          if (res.data.base.code == "10000") {
            console.log("获取信息成功");
            this.name = res.data.user.name;
            this.avatar = res.data.user.avatar;
            this.mail = res.data.user.email;
            this.signature = res.data.user.signature;
            this.CookiesSet = Cookies.set("avatar", this.avatar);
          }
          if (res.data.base.code == "10007") {
            console.log("不存在");
            this.name = "该用户不存在";
          }
          console.log(this.avatar);
          if (this.avatar === "") {
            this.avatar = "//xiey.work/ava.png";
          }
          if (this.signature === "") {
            this.signature = "这个人很懒，什么都没写。";
          }
        })
        .catch((err) => {
          console.error(err);
        });
    }
  },
  mounted() {
    this.id = Cookies.get("id");
    this.access_token = Cookies.get("access_token");
    this.refresh_token = Cookies.get("refresh_token");
    this.init();
  },
  components: {},
  computed: {},
  watch: {
    setcount(newsetcount) {
      if (newsetcount % 2 == 1) {
        this.setname = "设置";
      }
      if (newsetcount % 2 != 1) {
        this.setname = "关闭";
      }
    }
  },
  props: {},
  emits: {},
  setup() {}
};
</script>
<template>
  <a-page-header
    style="border: 1px solid rgb(235, 237, 240); background-color: #fff"
    title="个人中心"
    :sub-title="`id:` + id"
    @back="() => $router.go(-1)"
  />

  <div class="settings">
    <button class="button" @click="settings">
      <svg
        xmlns="https://www.w3.org/2000/svg"
        width="20"
        viewBox="0 0 20 20"
        height="20"
        fill="none"
        class="svg-icon"
      >
        <g stroke-width="1.5" stroke-linecap="round" stroke="#5d41de">
          <circle r="2.5" cy="10" cx="10"></circle>
          <path
            fill-rule="evenodd"
            d="m8.39079 2.80235c.53842-1.51424 2.67991-1.51424 3.21831-.00001.3392.95358 1.4284 1.40477 2.3425.97027 1.4514-.68995 2.9657.82427 2.2758 2.27575-.4345.91407.0166 2.00334.9702 2.34248 1.5143.53842 1.5143 2.67996 0 3.21836-.9536.3391-1.4047 1.4284-.9702 2.3425.6899 1.4514-.8244 2.9656-2.2758 2.2757-.9141-.4345-2.0033.0167-2.3425.9703-.5384 1.5142-2.67989 1.5142-3.21831 0-.33914-.9536-1.4284-1.4048-2.34247-.9703-1.45148.6899-2.96571-.8243-2.27575-2.2757.43449-.9141-.01669-2.0034-.97028-2.3425-1.51422-.5384-1.51422-2.67994.00001-3.21836.95358-.33914 1.40476-1.42841.97027-2.34248-.68996-1.45148.82427-2.9657 2.27575-2.27575.91407.4345 2.00333-.01669 2.34247-.97026z"
            clip-rule="evenodd"
          ></path>
        </g>
      </svg>
      <span class="lable">{{ setname }}</span>
    </button>
  </div>
  <div class="settingpage">
    <transition name="fade">
      <div class="setpage" v-show="setshow">
        <mdui-list :blur="true">
          <text>设置界面</text>
          <mdui-list-item @click="refreshtoken">刷新令牌</mdui-list-item>
          <mdui-list-item @click="logout">退出登陆</mdui-list-item>
        </mdui-list>
      </div>
    </transition>
    <div class="setpage-flur" v-show="setshow"></div>
  </div>
  <div class="upload">
    <transition name="fade">
      <div class="avatarpage" v-show="showavatarpage">
        <div class="up">
          <img
            v-if="previewImage"
            :src="previewImage"
            alt="预览图像"
            style="
              object-fit: cover;
              border-radius: 20px 20px 0 0;
              object-fit: cover;
              width: 100%;
              max-height: 250px;
            "
          />
          <form @submit.prevent="uploadAvatar">
            <a-input
              type="file"
              ref="avatarInput"
              @change="onFileChange"
              style="margin-top: 20px"
              accept="image/*"
              >上传文件</a-input
            >
          </form>
          <a-button type="primary submit" @click="putavatar(file)" style="margin-bottom: 20px"
            >上传头像</a-button
          >
        </div>
      </div>
    </transition>
    <div class="avatarpage-flur" v-show="showavatarpage" @click="modavatarpage"></div>
  </div>
  <div id="about">
    <div class="user_img">
      <img
        :src="avatar"
        class="touxiangimg rounded-image"
        @click="modavatarpage"
        style="object-fit: cover; width: 50vw; height: 50vw; max-width: 300px; max-height: 300px"
      />
    </div>
    <div class="user_name">
      <text id="name">{{ name }}</text>
    </div>
    <div>
      <el-button plain @click="toifollow">关注列表</el-button
      ><el-button plain @click="tofollowme">粉丝列表</el-button>
    </div>
    <!-- <input type="text" placeholder="请输入内容" class="input-box" /> -->
    <text>个人签名</text>

    <div class="signature">
      <div v-if="isEditing" class="input" style="z-index: 500">
        <input
          type="text"
          v-model="signature"
          @blur="savesignature()"
          autofocus
          style="z-index: 6665"
        />
        <el-button plain @click="savesignature()">确认</el-button>
      </div>
      <div v-else class="sign">
        <p @click="editsignature()">{{ signature }}</p>
      </div>
    </div>
    <div class="sign-flur" v-show="isEditing" @click="savesignature()"></div>

    <text>邮箱:{{ mail }}</text>
  </div>

  <a-divider orientation="left" class="separate">评价</a-divider>

  <br />

  <!-- <div class="bg"></div> -->
</template>

<style scoped>
#name {
  position: relative;
  font-size: 2em; /* 设置字体大小为1.5em */
  font-weight: bold; /* 设置字体粗细为粗体 */
  margin-block: 0.5em; /* 设置行距为0.5em */
}
.rounded-image {
  border-radius: 50%; /* 半径为50%意味着图片将变成圆形 */
}
#about {
  display: grid;
  justify-items: center;
  grid-template-columns: repeat(1, 1fr);
  gap: 10px 10px;
  margin-top: 50px;
}
.sign {
  display: inline-block;
  padding: 5px 5px 5px;
  border: 2px dashed #ccc;
  border-radius: 10px;
}

.avatarpage-flur {
  opacity: 1;
  position: fixed;
  width: 100%;
  height: 100%;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  z-index: 100;
  backdrop-filter: saturate(100%) blur(5px);
  background: inherit;
}
.avatarpage {
  position: fixed;

  gap: 10px;
  width: 90%;
  height: auto;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  border-radius: 20px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.2); /* 添加阴影效果 */
  z-index: 9999;
  max-width: 500px;
}
.up {
  width: 100%;
  margin-left: 5%;
  margin-top: 5%;
  display: grid;
  justify-items: center;
  align-items: center;
  gap: 10px;
  height: 100%;
  margin-bottom: 5%;
  margin-block: 0;
  margin-inline: 0;
  padding-block: 0;
  padding-inline: 0;
  padding: 0;
}

.bg {
  background-image: url("https://saas.bk-cdn.com/t/f1e16078-bb12-419d-ad3f-be967f268d40/u/f0896296-9cb3-4e7a-bfaa-44a0c8c681f0/1715359885684/117612906_p0%20%281%29.jpg");
  width: 100%;
  height: 100%;
  position: fixed;
  top: 0;
  left: 0;
  background-repeat: no-repeat;
  background-attachment: fixed;
  background-size: cover;
  background-position: center;
  pointer-events: none;
  z-index: -1;
  opacity: 0.35;
}

.setpage {
  position: fixed;
  top: 150px;
  right: 20px;
  height: 60%;

  width: 150px;
  border-radius: 15px;
  align-items: center;
  display: flex;
  flex-direction: column;
  gap: 10px;
  padding: 3px 3px 3px 3px; /* 上下左右内边距 */
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.2); /* 添加阴影效果 */
  z-index: 1060; /* 确保设置界面在其它内容之上 */
}

.setpage-flur {
  opacity: 1;
  position: fixed;
  top: 150px;
  right: 20px;
  height: 60%;
  width: 150px;
  z-index: 1050;
  backdrop-filter: saturate(100%) blur(10px);
  background: inherit;
}
.setpage text {
  position: relative;
  left: 22.5px;
  font-weight: bold;
}
.setpage mdui-list-item {
  margin-left: 5px;
}
.fade-enter {
  transition: opacity 0.2s;
}

.fade-enter-active {
  opacity: 1;
}

.fade-leave-active {
  transition: opacity 0.2s;
}

.fade-leave-to {
  opacity: 0;
}

.settings {
  position: fixed;
  right: 25px;
  top: 100px;
}

.button {
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 6px 12px;
  gap: 8px;
  height: 36px;
  width: 120px;
  border: none;
  background: #5e41de33;
  border-radius: 20px;
  cursor: pointer;
}

.lable {
  line-height: 20px;
  font-size: 17px;
  color: #5d41de;
  font-family: sans-serif;
  letter-spacing: 1px;
}

.button:hover {
  background: #5e41de4d;
}

.button:hover .svg-icon {
  animation: spin 2s linear infinite;
}

@keyframes spin {
  0% {
    transform: rotate(0deg);
  }

  100% {
    transform: rotate(360deg);
  }
}
.separate {
  margin-top: 30px;
  font-size: 18px;
}
</style>
