# 本项目前端使用Vue+element UI 开发,本着先学会用,再学原理的步骤,这里记录其中学到的一些语法用法
## 本项目为从零开始
---
1. 缩写: @click="" 表示v-on:click=""  表示监听事件
2. 缩写: :src="" 表示v-bind:src=""  表示单向绑定
3. 双向绑定: v-model="" 常用于form表单输入
  ```html
  <el-form :model="form" @submit.prevent="login">
    <el-form-item label="用户名" :label-width="formLabelWidth">
      <el-input v-model="form.username" autocomplete="off" clearable></el-input>
    </el-form-item>
    <el-form-item label="密码" :label-width="formLabelWidth">
      <el-input v-model="form.password" autocomplete="off" show-password></el-input>
    </el-form-item>
  </el-form>
  <!-- 这里的el-form的属性model="form"是可以去掉的 -->
  ```
  ```html
  <input v-model="value" />
  <!-- 相当于 -->
  <input type="text" 
  　　　:value="value" 
  　　　@input="value=$event.target.value" />
  <input
  :value="text"
  @input="e => text = e.target.value"
  />
  ```
4. vue-resourse `this.$http.get(url)`
```html
<script src="https://cdn.staticfile.org/vue-resource/1.5.1/vue-resource.min.js"></script>
注意是异步加载 
```
5. 回到顶部
```html
<el-backtop target=".test" right="120">UP</el-backtop>
给予内容对象 class="test"即可
```
