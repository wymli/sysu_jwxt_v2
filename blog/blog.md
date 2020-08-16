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
6. 标签页切换事件
   1. v-model 双向绑定  然后watch
   2. @tab-click 监听事件 handleClick(tab, event)
```html
<el-tab v-model="activeName">
  <el-tab-pane name="first"></el-tab-pane>
  <el-tab-pane name="second"></el-tab-pane>
  <el-tab-pane name="third"></el-tab-pane>
</el-tabs>

new Vue({
  watch:{
    "activeName" : function(val){
      //这里val = name
    }
  }
})
```
7. 按钮大小自动适应内部文字
```html
<el-button style="width:fit-content;height:fit-content;" />
```
8. 按钮放置于屏幕中央
```html
<el-button style="width:fit-content;height:fit-content;position:absolute;top:0;left:0;right:0;bottom:0;margin:auto auto;" />
<!-- 思路就是外边距边界为整个屏幕,然后调整外边距
所以这个的父元素要 `width:100% height:100%` -->
```
9. 子元素在父元素中的左右浮动
```html
子元素设置float:right; 然后通过外边距边界微调:right/left:5%
有时候right/left不行(不知道为什么),可以使用margin-left/margin-right,
本质都是调整border位置
```