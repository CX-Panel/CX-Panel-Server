<style scoped lang="less">
    .login-warp {
        width: 350px;
        margin: 0px auto !important;
        padding-top: 100px;
    }
    .card-warp {
        background-color: #fff;
    } 
    .login-submit-warp {
        text-align: center;
    }
</style>

<template>
    <div class="login-warp">
        <Card dis-hover class="card-warp">
            <Form ref="formInline" :model="formInline" :rules="ruleInline" >
                <h1 class="text-center mb-3">用户登录</h1>
                <FormItem prop="username">
                    <Input type="text"  v-model="formInline.username" placeholder="用户名">
                        <Icon type="ios-person-outline" slot="prepend"></Icon>
                    </Input>
                </FormItem>
                <FormItem prop="password">
                    <Input type="password"  v-model="formInline.password" placeholder="密码">
                        <Icon type="ios-lock-outline" slot="prepend"></Icon>
                    </Input>
                </FormItem>
                <FormItem class="login-submit-warp">
                    <Button type="primary" size='large' @click="handleSubmit('formInline')">登录</Button>
                </FormItem>
            </Form>
        </Card>
    </div>
</template>

<script>
    export default {
        data () {
            return {
                formInline: {
                    username: '',
                    password: ''
                },
                ruleInline: {
                    username: [
                        { required: true, message: '请输入用户名', trigger: 'blur' }
                    ],
                    password: [
                        { required: true, message: '请输入密码', trigger: 'blur' },
                    ]
                }
            }
        },
        methods: {
            handleSubmit(name) {
                this.$refs[name].validate((valid) => {
                    if (valid) {
                        this.$Message.success('Success!');
                    } else {
                        this.$Message.error('Fail!');
                    }
                })
            }
        }
    }
</script>