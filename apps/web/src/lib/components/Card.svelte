<script lang="ts">
	import Link from './Link.svelte';

	interface CardProps {
		thumbnail: {
			src: string;
			alt: string;
		};
		title: string;
		description: string;
		link?: {
			href: string;
			label: string;
		};
		variant: 'primary' | 'secondary' | 'tertiary';
	}

	let { thumbnail, title, description, link, variant = 'primary' }: CardProps = $props();
</script>

<div class="flex h-full grow flex-col gap-5.5">
	<img src={thumbnail.src} alt={thumbnail.alt} class="rounded-sm" />

	<h3 class="text-3xl font-normal">{title}</h3>

	<p>{description}</p>

	{#if link}
		<Link href={link.href} {variant} cn="mt-auto">{link.label}</Link>
	{/if}
</div>
