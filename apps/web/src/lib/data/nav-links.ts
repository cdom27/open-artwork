export const headerLinks = [{ name: 'Gallery', href: '/gallery', subLinks: null }];

export const footerLinks = [
	{
		name: 'Website',
		links: [
			{ name: 'Home', href: '/' },
			{ name: 'API', href: '/docs', comingSoon: true },
			...headerLinks,
			{ name: 'Timeline', href: '/timeline', comingSoon: true },
			{ name: 'Color Trends', href: '/color-trends', comingSoon: true }
		]
	},
	{
		name: 'Support',
		links: [
			{ name: 'Star on GitHub', href: 'https://github.com/cdom27/open-artwork' },
			{ name: 'Report a bug', href: '#bug-report', comingSoon: true },
			{ name: 'Report a data issue', href: '#data-issue', comingSoon: true }
		]
	},
	{
		name: 'Contact',
		links: [
			{ name: 'Hello@openart.work', href: 'mailto:hello@openart.work' },
			{ name: 'Contact Form', href: '#general-contact', comingSoon: true }
		]
	}
];
