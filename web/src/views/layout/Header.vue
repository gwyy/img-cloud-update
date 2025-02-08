<template>
    <div class="header">
       <div class="header-left">
            <router-link class="header-left-item" to="/index">云上传</router-link>
       </div>
       <div class="header-right">
            <ul>
                <li class="header-right-item">
                    <router-link to="/index">首页</router-link>
                </li>
                <li class="header-right-item">
                    <router-link to="/setting">设置</router-link>
                </li>
            </ul>
       </div>
    </div>
</template>
<style scoped>
.header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    height: 60px;
    padding: 0 50px;
    background: linear-gradient(120deg, #2d6cdf 0%, #4c89f0 50%, #6ea5ff 100%);
    box-shadow: 0 2px 10px rgba(0,0,0,0.15);
}

.header-left {
    display: flex;
    align-items: center;
}

.header-left-item {
    font-size: 20px;
    font-weight: 500;
    color: #fff;
    text-decoration: none;
    transition: all 0.3s;
}

.header-left-item:hover {
    color: #e6f0ff;
}

.header-right ul {
    display: flex;
    list-style: none;
    margin: 0;
    padding: 0;
}

.header-right-item {
    margin-right: 30px;
}

.header-right-item a {
    font-size: 15px;
    color: #e6f0ff;
    text-decoration: none;
    transition: all 0.3s;
    padding: 8px 0;
    position: relative;
}

.header-right-item a::after {
    content: '';
    position: absolute;
    bottom: 0;
    left: 0;
    width: 0;
    height: 2px;
    background: #fff;
    transition: all 0.3s;
}

.header-right-item a:hover {
    color: #fff;
}

.header-right-item a:hover::after {
    width: 100%;
}

.header-right-item:last-child {
    margin-right: 0;
}
</style>
