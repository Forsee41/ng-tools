import { defineConfig } from "vite";

export default defineConfig({
	root: "./frontend",
	build: {
		rollupOptions: {
			input: {
				main: "./frontend/index.html",
			},
		},
	},
	server: {
		port: 3000,
		proxy: {
			"/api": {
				target: "http://localhost:8080",
				changeOrigin: true,
			},
		},
	},
});
