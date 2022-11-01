<template>
<div class="login">
    <el-container>
        <el-main>
    <el-form :model="user" :rules="rules">
        <el-form-item label="Telephone" prop="telephone">
            <el-input v-model="user.telephone" />
        </el-form-item>
        <el-form-item label="Password" prop="password">
            <el-input type="password" show-password v-model="user.password"/>
        </el-form-item>
        <el-form-item>
            <el-button type="primary" @click="login">Login</el-button>
        </el-form-item>
    </el-form>
    </el-main>
    </el-container>
</div>
</template>

<script>
import { mapActions } from 'vuex';
import { ElMessageBox } from 'element-plus';

export default {

  data() {
    return {
      user: {
        telephone: '',
        password: '',
      },
      rules: {
        telephone: [
          { required: true, message: 'please input telephone', trigger: 'blur' },
          { pattern: /^1[3|4|5|7]\d{9}$/, message: '电话不合法', trigger: 'blur' },
        ],
        password: [
          { required: true, message: 'please input password', trigger: 'blur' },
          { min: 6, message: '密码不得少于六位', trigger: 'blur' },
        ],
      },
    };
  },
  methods: {
    // ...mapMutations('userModule', ['SET_TOKEN', 'SET_USERINFO']),
    ...mapActions('userModule', { userLogin: 'login' }),
    login() {
      this.userLogin(this.user).then(() => {
        this.$router.replace({ name: 'home' });
      }).catch((err) => {
        this.errTip();
        console.log(err);
      });
    },
    errTip() {
      ElMessageBox.alert('用户已存在', '错误提示');
    },
  },
};
</script>

<style lang="scss" scoped>

</style>
