<script lang="ts">
  import {CheckBuildNumber, CheckLogin} from "../wailsjs/go/main/App";
  import Dashboard from "./components/Dashboard.svelte";
  import LoadingDots from "./components/layovers/LoadingDots.svelte";
  import Toasts from "./components/alerts/Toasts.svelte";
  import {addToast} from "./components/alerts/ts/store";

  let user: string = '';
  let password: string = '';
  let loading: boolean = false
  let loginState: boolean = true

  async function HandleLogin() {
    loading = true
    if (password.length !== 4) {
      addToast({
        message:"Invalid login",
        type: "error",
        dismissible: true,
        timeout: 10000
      })
      return
    }
    try{
      await CheckLogin(user, password).then((async result => {
        if (!result) {
          addToast({
            message:"Invalid login",
            type: "error",
            dismissible: true,
            timeout: 10000
          })
          setNewState(true)
          loading = false
          return
        } else {
          setNewState(false)
        }
      }))
      await CheckBuildNumber().then(r => {
        if (r < 0) {
          addToast({
            message:"Your app is up to date",
            type: "success",
            dismissible: true,
            timeout: 3500
          })
        } else {
          addToast({
            message:"There is an update available, build number: " + r.toString(),
            type: "info",
            dismissible: true,
            timeout: 10000
          })
        }
      })
    } catch (e) {
      loading = false
      console.log(e)
    }
    loading = false
  }

  function setNewState(state: boolean) {
    loginState = state
  }

</script>

<Toasts />

<main>
  {#if loginState}
    <div class="login-container">
      <div class="login-box">
        <h2 style="color: black">Login</h2>
        <input bind:value={user} type="email" placeholder="User" />
        <input bind:value={password} type="password" placeholder="Pin" />
        <button on:click={async ()=>{await HandleLogin().then(r => {})}}>Login</button>
        <!--{#if err !== undefined}-->
        <!--  <div style="color: crimson">{err}</div>-->
        <!--{/if}-->
      </div>
      {#if loading}
        <div style="position: fixed; top: 95%; left: 50%">
          <LoadingDots />
        </div>
      {/if}
    </div>
  {/if}
  {#if !loginState}
    <Dashboard />
  {/if}
</main>

<style>
  .login-container {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
    font-family: 'Roboto', sans-serif;
    /*background-color: #f2f2f2;*/
  }

  .login-box {
    width: 300px;
    padding: 20px;
    border-radius: 10px;
    background-color: #fff;
    box-shadow: 0 0 10px rgba(0,0,0,0.1);
    box-sizing: border-box;
  }

  .login-box h2 {
    margin-bottom: 20px;
    text-align: center;
    color: #333;
    /*font-weight: 500;*/
  }

  .login-box input {
    width: 100%;
    padding: 10px;
    margin-bottom: 10px;
    border: 1px solid #ddd;
    border-radius: 5px;
    font-size: 16px;
    box-sizing: border-box;
  }

  .login-box button {
    width: 100%;
    padding: 10px;
    border: none;
    border-radius: 5px;
    background-color: #007BFF;
    color: #fff;
    font-size: 16px;
    cursor: pointer;
    transition: background-color 0.2s ease;
  }

  .login-box button:hover {
    background-color: #0056b3;
  }
</style>
