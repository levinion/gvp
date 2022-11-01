<template>
    <div class="register">
           <el-container>
        <el-main>
    <el-form :model="user" :rules="rules">
        <el-form-item label="UserName" prop="name">
            <el-input v-model="user.name"/>
        </el-form-item>
        <el-form-item label="Telephone" prop="telephone">
        <el-input v-model="user.telephone"/>
        </el-form-item>
        <el-form-item label="Password" prop="password">
            <el-input type="password" show-password v-model="user.password"/>
        </el-form-item>
        <el-form-item>
            <el-button type="primary" @click="register">register</el-button>
        </el-form-item>
    </el-form>
    </el-main>
    </el-container>
    </div>
</template>

<script>
import { ElMessageBox } from 'element-plus';
import { mapActions } from 'vuex';
// import { mapMutations } from 'vuex';

export default {
  data() {
    return {
      user: {
        name: '',
        password: '',
        elephone: '',
      },
      rules: {
        name: [
          { required: false, message: 'please input username', trigger: 'blur' },
        ],
        telephone: [
          { required: true, message: 'please input telephone', trigger: 'blur' },
          {
            pattern: /^1[3|4|5|7]\d{9}$/, message: '电话号码不合法', trigger: ['blur'],
          },
        ],
        password: [
          { required: true, message: 'please input password', trigger: 'blur' },
          { min: 6, message: '密码不得少于六位', trigger: ['blur'] },
        ],
      },
    };
  },
  methods: {
    // ...mapMutations('userModule', ['SET_TOKEN', 'SET_USERINFO']),
    ...mapActions('userModule', { userRegister: 'register' }),
    register() {
      this.userRegister(this.user).then(() => {
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
