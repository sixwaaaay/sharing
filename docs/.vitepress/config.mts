import { defineConfig } from 'vitepress'

export default defineConfig({
  title: "sharing",
  description: "document for sharing",
  themeConfig: {
    nav: [
      { text: 'Home', link: '/' },
    ],
    sidebar: [
      {
        text: "comments",
        base: "/comments/",
        items: [
          { text: 'Overview', link: '/index' },
        ]
      },
      {
        text: "users",
        base: "/users/",
        items: [
          { text: 'Overview' , link: '/index' },
        ]
      },
      {
        text: "Content",
        base: "/content/",
        items: [
          { text: 'Overview' , link: '/index' },
        ]
      },
      {
        text: "Graph",
        items: [
          { text: 'Overview' },
        ]
      },
      {
        text: "CDC",
        items: [
          { text: 'Overview' },
        ]
      },
      {
        text: "Cron",
        items: [
          { text: 'Overview' },
        ]
      },
    ],

    socialLinks: [
      { icon: 'github', link: 'https://github.com/sixwaaaay/sharing' }
    ]
  }
})
