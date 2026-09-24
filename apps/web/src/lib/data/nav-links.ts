export const headerLinks = [
	{ name: 'Gallery', href: '/gallery', subLinks: null },
	{ name: 'Timeline', href: '/timeline' },
	{ name: 'Color Trends', href: '/color-trends' },
	{ name: 'API', href: '/docs', subLinks: null },
	{ name: 'About', href: '/about', subLinks: null }
];

export const footerLinks = [
	{
		name: 'Website',
		links: [{ name: 'Home', href: '/' }, ...headerLinks]
	},
	{
		name: 'Support',
		links: [
			{ name: 'Star on GitHub', href: 'https://github.com/cdom27/open-artwork' },
			{ name: 'Report a bug', href: '#bug-report' },
			{ name: 'Report a data issue', href: '#data-issue' }
		]
	},
	{
		name: 'Contact',
		links: [
			{ name: 'Hello@openart.work', href: 'mailto:hello@openart.work' },
			{ name: 'Message Form', href: '#general-contact' }
		]
	}
];
