<script lang="ts">
	import type { Snippet } from 'svelte';
	import type { HTMLAnchorAttributes } from 'svelte/elements';
	import { variants } from '$lib/utils/link-button-variants';

	interface AnchorProps extends HTMLAnchorAttributes {
		variant: 'primary' | 'secondary' | 'tertiary';
		children: Snippet<[]>;
		cn?: string;
	}

	let { variant = 'primary', children, cn = '', ...rest }: AnchorProps = $props();

	const variantClass = variants[variant];
	const classes =
		`${variantClass} transition-all duration-300 cursor-pointer rounded-sm py-1.5 px-2.5 hover:rounded-lg active:rounded-xl ${cn}`.trim();
</script>

<a class={classes} {...rest}>{@render children()}</a>
