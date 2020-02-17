<style scoped lang="less">
   
</style>
<template>
    <div>
        <header class="cx-header">
            <div class="cx-header-logo"><a href="" class="cx-header-logo-link">面板管理</a></div>
            <!-- <ul class="cx-header-menu">
                <li>
                    <a href="" class="cx-header-menu-link">导航</a>
                </li>
            </ul> -->
        </header>
        <aside class="cx-aside">
            <div class="cx-menu-wrap">
                <div class="cx-menu">
                    <ul class="cx-menu-list">
                        <li class="cx-menu-item">
                            <a href="" class="cx-menu-link active">站点管理</a>
                        </li>
                        <li class="cx-menu-item">
                            <a href="" class="cx-menu-link">数据库管理</a>
                        </li>
                        <li class="cx-menu-item">
                            <a href="" class="cx-menu-link">FTP管理</a>
                        </li>
                        <li class="cx-menu-item">
                            <a href="" class="cx-menu-link">文件管理</a>
                        </li>
                        <li class="cx-menu-item">
                            <a href="" class="cx-menu-link">应用中心</a>
                        </li>
                        <li class="cx-menu-item">
                            <a href="" class="cx-menu-link">软件管理</a>
                        </li>
                    </ul>
                </div>
            </div>
        </aside>
        <main class="cx-main-container">
            <header class="admin-container-header">
                <div class="header-group">
                    <h4 class="title">站点</h4>
                    <h5 class="sub-title">全部站点</h5>
                </div>
            </header>
            <div class="admin-container-body">
                <div class="panel-box">
                    <div class="panel-box-body">
                        <!--内容列表-->
                        <section class="diy-table-list" v-cloak>
                            <div class="diy-table-item diy-table-item-header">
                                <ul class="list-unstyled row">
                                    <li class="col-md-1">
                                        <Checkbox v-model="itemListSelectStatus"  :disabled='!itemList.length' @on-change='itemListSelectChange'>全选</Checkbox>
                                    </li>
                                    <li class="col-md-1">
                                        分类
                                    </li>
                                    <li class="col-md-3"><span>名称</span></li>
                                    <li class="col-md-2"><span>时间</span></li>
                                    <li class="col-md-1"><span>浏览</span></li>
                                    <li class="col-md-2"><span>操作</span></li>
                                </ul>
                            </div>
                            <div v-if='itemList.length' class="diy-table-body">
                                <div class="diy-table-item" v-for='item in itemList' >
                                     <ul class="list-unstyled row">
                                        <li class="col-md-1"><Checkbox v-model="item.itemListSelectStatus"></Checkbox></li>
                                        <li class="col-md-1 over-h-e"> </li>
                                        <li class="col-md-3 over-h-e"><span><a :href="item.url" target="_blank">{{item.title}}</a></span></li>
                                        <li class="col-md-2"><span>{{item.publish_time | time}}</span></li>
                                        <li class="col-md-1"><span>{{item.views}}</span></li>
                                        <li class="col-md-2">
                                            <span>
                                                <a @click="itemDel(item)" >删除</a>   
                                                <a @click="itemEdit(item)" class="ml-1">编辑</a>
                                            </span>
                                        </li>
                                        
                                    </ul>
                                </div>
                            </div>
                            <div class="no-block" v-else>
                                <Icon type="ios-filing-outline"></Icon>
                                <p>暂无数据</p>
                            </div>
                        </section>
                    </div>
                </div>
            </div>
        </main>
    </div>
        
        
</template>
<script>
    export default {
        data () {
            return {
               itemListSelectStatus: false,
               itemList: [],
            }
        },
        methods: {
            itemDel() {},
            itemEdit() {},
            itemListSelectChange(){},
        }
    }
</script>
