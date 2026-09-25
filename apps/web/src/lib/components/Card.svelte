<script lang="ts">
	import type { HTMLImgAttributes } from 'svelte/elements';
	import Link from './Link.svelte';

	interface CardProps {
		thumbnail: HTMLImgAttributes;
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
	<img
		{...thumbnail}
		class="rounded-sm opacity-0 transition-all duration-400 group-[&.in-view]:opacity-100"
	/>

	<h3
		class="translate-y-5 text-3xl font-normal opacity-0 transition-all delay-160 duration-400 group-[&.in-view]:translate-y-0 group-[&.in-view]:opacity-100"
	>
		{title}
	</h3>

	<p
		class="translate-y-5 opacity-0 transition-all delay-240 duration-400 group-[&.in-view]:translate-y-0 group-[&.in-view]:opacity-100"
	>
		{description}
	</p>

	{#if link}
		<div
			class="mt-auto flex translate-y-5 flex-col gap-2.5 opacity-0 transition-all delay-240 duration-400 group-[&.in-view]:translate-y-0 group-[&.in-view]:opacity-100 sm:flex-row sm:gap-4"
		>
			<Link href={link.href} {variant}>{link.label}</Link>
		</div>
	{/if}
</div>
