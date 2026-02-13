import { defineConfig } from 'vitepress'
import {
	groupIconMdPlugin,
	groupIconVitePlugin,
} from 'vitepress-plugin-group-icons'
import llmstxt, { copyOrDownloadAsMarkdownButtons } from 'vitepress-plugin-llms'

const version = 'latest'

export default defineConfig({
	base: '/gostandards/',
	title: 'gostandards',
	description: 'Collection of standards as types',
	themeConfig: {
		logo: '/logo.png',
		outline: [2, 4],
		nav: [
			{
				text: 'Getting Started',
				link: '/guide/getting-started',
			},
			{
				text: `${version}`,
				items: [
					{
						text: 'Release Notes',
						link: 'https://github.com/foomo/gostandards/releases',
					},
				],
			},
		],
		sidebar: [
			{
				text: 'Guide',
				items: [
					{ text: 'Introduction', link: '/' },
					{ text: 'Getting Started', link: '/guide/getting-started' },
				],
			},
			{
				text: 'Standards',
				items: [
					{ text: 'ISO 3166 - Country Codes', link: '/standards/iso3166' },
					{ text: 'ISO 4217 - Currency Codes', link: '/standards/iso4217' },
					{ text: 'HTTP', link: '/standards/http' },
				],
			},
			{
				text: 'Contributing',
				items: [
					{
						text: "Guideline",
						link: '/CONTRIBUTING.md',
					},
					{
						text: "Code of conduct",
						link: '/CODE_OF_CONDUCT.md',
					},
					{
						text: "Security guidelines",
						link: '/SECURITY.md',
					},
				],
			},
		],
		editLink: {
			pattern: 'https://github.com/foomo/gostandards/edit/main/docs/:path',
			text: 'Suggest changes to this page',
		},
		search: {
			provider: 'local',
		},
		footer: {
			message: 'Made with ♥ <a href="https://www.foomo.org">foomo</a> by <a href="https://www.bestbytes.com">bestbytes</a>',
		},
		socialLinks: [
			{
				icon: 'docker',
				link: 'https://hub.docker.com/u/foomo',
			},
			{
				icon: 'github',
				link: 'https://github.com/foomo/gostandards',
			},
		],
	},
	head: [
		['meta', { name: 'theme-color', content: '#ffffff' }],
		['link', { rel: 'icon', href: '/logo.png' }],
		['meta', { name: 'author', content: 'foomo by bestbytes' }],
		['meta', { property: 'og:title', content: 'foomo/gostandards' }],
		[
			'meta',
			{
				property: 'og:image',
				content: 'https://github.com/foomo/gostandards/blob/main/docs/public/banner.png?raw=true',
			},
		],
		[
			'meta',
			{
				property: 'og:description',
				content: 'Collection of standards as types',
			},
		],
		['meta', { name: 'twitter:card', content: 'summary_large_image' }],
		[
			'meta',
			{
				name: 'twitter:image',
				content: 'https://github.com/foomo/gostandards/blob/main/docs/public/banner.png?raw=true',
			},
		],
		[
			'meta',
			{
				name: 'viewport',
				content: 'width=device-width, initial-scale=1.0, viewport-fit=cover',
			},
		],
	],
	markdown: {
		theme: {
			dark: 'one-dark-pro',
			light: 'github-light',
		},
		config(md) {
			md.use(groupIconMdPlugin)
			md.use(copyOrDownloadAsMarkdownButtons)
		},
	},
	vite: {
		plugins: [
			groupIconVitePlugin(),
			llmstxt({
				excludeIndexPage: false,
			}),
		],
	},
	sitemap: {
		hostname: 'https://foomo.github.io/gostandards',
	},
	ignoreDeadLinks: true,
})
