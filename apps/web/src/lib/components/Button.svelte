<script lang="ts">
	import type { Snippet } from 'svelte';
	import type { HTMLButtonAttributes } from 'svelte/elements';
	import { variants } from '$lib/utils/link-button-variants';

	interface ButtonProps extends HTMLButtonAttributes {
		variant: 'primary' | 'secondary' | 'tertiary' | 'outline';
		children: Snippet<[]>;
		cn?: string;
	}

	let { variant = 'primary', children, cn = '', ...rest }: ButtonProps = $props();

	const variantClass = $derived(variants[variant]);
	const classes = $derived(
		`${variantClass} transition-all duration-300 cursor-pointer rounded-sm py-1.5 px-2.5 hover:rounded-lg active:rounded-xl self-start ${cn}`.trim()
	);
</script>

<button class={classes} {...rest}>{@render children()}</button>
