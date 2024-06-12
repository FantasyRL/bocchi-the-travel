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
      file: null
    };
  },
  methods: {
    follow(type) {
      axios
        .post(
          this.url +
            "https://api.xiey.work/bocchi/trust/action?object_uid=" +
            this.id +
            "&action_type=" +
            type,
          {},
          {
            headers: {
              "access-token": this.access_token
            }
          }
        )
        .then((res) => {
          console.log(res.data.base.code);
          alert(res.data.base.msg);
        });
    },
    init() {
      axios
        .get(this.url + "/bocchi/user/info?user_id=" + this.id, {
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
    this.id = Number(this.$route.params.id);
    this.access_token = Cookies.get("access_token");
    this.refresh_token = Cookies.get("refresh_token");
    this.init();
  },
  components: {},
  computed: {},
  watch: {},
  props: {},
  emits: {},
  setup() {}
};
</script>
<template>
  <div class="about">
    <a-page-header
      style="border: 1px solid rgb(235, 237, 240); background-color: #fff"
      title="Ta的资料"
      sub-title=""
      @back="() => $router.go(-1)"
    />

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

      <text>个人签名</text>
      <div class="signature"></div>
      <div class="sign-flur" v-show="isEditing" @click="savesignature(this.signature)"></div>
      <text>uid:{{ id }}</text>
      <text>邮箱:{{ mail }}</text>
      <div class="follow">
        <el-button type="primary" @click="follow(1)">关注</el-button>
        <el-button type="danger" @click="follow(0)">取消关注</el-button>
      </div>
    </div>
    <a-divider orientation="left" class="separate">评价</a-divider>
  </div>

  <br />
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
.rounded-image {
  border-radius: 50%; /* 半径为50%意味着图片将变成圆形 */
}
.touxiangimg {
  height: 200px;
  width: 200px;

  margin-block: 0;
  max-width: 200px;
}

.avatarpage-flur {
  opacity: 1;
  position: fixed;
  width: 100%;
  height: 100%;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  z-index: 1050;
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
.about {
  z-index: 5;
  position: relative;
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
.info {
  z-index: 0;
}
.touxiang {
  z-index: inherit;
  display: grid;
  justify-items: center;
}
.shejiao {
  /* 社交账号 */
  display: flex;
  flex-direction: column;
  align-items: center;
  background-color: 000000;
  position: relative;
  height: auto;
  width: 90%;
  margin: auto;
  text-align: center;
  z-index: 5;
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
/* .fade-enter-active, .fade-leave-active {
  transition: opacity .5s;
}
.fade-enter, .fade-leave-to
  opacity: 0;
} */
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
