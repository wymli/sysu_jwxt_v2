package main

import (
	"github.com/gin-gonic/gin"
)

func (client *Client) ifNotLoginAndReturn(ctx *gin.Context) bool {
	if client == nil || client.isLogin == false {
		ctx.Header("Content-Type", "text/html;charset=utf-8")
		ctx.String(404, `
		<html>
		<link rel="stylesheet" href="https://unpkg.com/element-ui/lib/theme-chalk/index.css">
		<script src="https://unpkg.com/vue/dist/vue.js"></script>
		<script src="https://unpkg.com/element-ui/lib/index.js"></script>
		<p id="app">.</p>
		<style>
		body {
			margin: 0;
			padding: 0;
			width: 100%;
			height: 100%;
			background: url(/view/pic/background1.jpg) no-repeat;
			background-position: center;
			background-size: cover;
		}
		</style>
		<script>
		new Vue({
			created(){
				this.$message("小坏蛋,先登陆哦")
				setTimeout(() => {
					// window.location.replace("/")
					this.$message("Time Clocked")
				}, 4000)
			}
			
		})
		</script>
		</html>
		`)
		return true
	}
	return false
}
