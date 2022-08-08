<template>
  <div>
    <main class="cover w-screen min-h-screen bg-cover">
      <div class="flex">
        <div class="flex basis-1/2 justify-center items-center my-[78px]">
          <div
            class="box flex py-4 w-[500px] h-[700px] justify-center items-center rounded-3xl"
          >
            <div>
              <h2 class="header">Login</h2>
              <h3 class="header-small py-2">Enjoy your mood!</h3>
              <div class="register-form mt-12 space-y-6">
                <div class="flex flex-col space-y-4">
                  <label class="register-label" for="username">Username</label>
                  <input
                    class="register-input bg-transparent placeholder:opacity-50 border-b-2 py-1 border-gray-500 focus:outline-none focus:shadow-outline"
                    type="text"
                    placeholder="Enter your username"
                    v-model="form.username"
                  />
                </div>
                <div class="flex flex-col space-y-4">
                  <label class="register-label" for="password">Password</label>
                  <input
                    class="register-input bg-transparent placeholder:opacity-50 border-b-2 py-1 border-gray-500 focus:outline-none focus:shadow-outline"
                    type="password"
                    placeholder="Enter your password"
                    v-model="form.password"
                  />
                </div>
              </div>
              <button
                type="button"
                class="register-button my-4 w-full py-2 rounded-lg text-white font-medium"
                @click="clickLogin"
              >
                Login
              </button>
              <div class="text-center">
                <span class="font-medium"> Don't have an account ? </span>
                <router-link class="text-[#f96632] font-medium" to="/register"
                  >Register</router-link
                >
              </div>
            </div>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup>
import { reactive } from "vue";
import { fetchPost } from "../lib/api";
import { useRouter } from "vue-router";

const router = useRouter();
const form = reactive({
  username: "",
  password: "",
});

async function clickLogin() {
  const url = process.env.VUE_APP_URL + "login";
  const data = {
    username: form.username,
    password: form.password,
  };
  const response = await fetchPost(url, data, "");
  const results = response.data;

  if (response.status === 200) {
    if (results.message === "success") {
      window.localStorage.setItem("username", data.username);
      window.localStorage.setItem("token", results.token);
      router.push("/");
    } else {
      alert("Username or Password Incorrect.");
    }
  }
}
</script>

<style scoped>
.cover {
  background-image: url("@/assets/register_bg.png");
  background-repeat: no-repeat;
}

.box {
  background: rgba(255, 255, 255, 0.4);
  box-shadow: 10px 10px 10px rgba(0, 0, 0, 0.25);
}

.header {
  font-style: normal;
  font-weight: 600;
  font-size: 55px;
  line-height: 67px;
  text-align: center;

  color: #ffffff;

  text-shadow: 4px 4px 4px rgba(0, 0, 0, 0.25);
}

.header-small {
  font-style: normal;
  font-weight: 500;
  font-size: 20px;
  line-height: 24px;
  text-align: center;

  color: #000000;

  text-shadow: 0px 4px 4px rgba(0, 0, 0, 0.25);
}

.register-label {
  font-style: normal;
  font-weight: 500;
  font-size: 20px;
  line-height: 24px;
  color: #f96632;
}

.register-input {
  font-style: normal;
  font-weight: 500;
  font-size: 20px;
  line-height: 24px;

  color: rgba(0, 0, 0, 0.8);
}

.register-button {
  background: #ff914d;
  box-shadow: 2px 2px 2px #6c77c1;
}
</style>
