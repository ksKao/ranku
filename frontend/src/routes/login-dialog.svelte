<script lang="ts">
	import Button from '$lib/components/ui/button/button.svelte';
	import * as Dialog from '$lib/components/ui/dialog/index';
	import * as Tabs from '$lib/components/ui/tabs/index';
	import type { ComponentProps } from 'svelte';
	import LoginForm from './login-form.svelte';

	type Props = Pick<ComponentProps<typeof Dialog.Root>, 'open' | 'onOpenChange'> & {
		loginForm: ComponentProps<typeof LoginForm>['form'];
	};

	const { loginForm, ...props }: Props = $props();
</script>

<Dialog.Root {...props}>
	<Dialog.Content>
		<Dialog.Header>
			<Dialog.Title>Login</Dialog.Title>
			<Dialog.Description>Login to your account</Dialog.Description>
		</Dialog.Header>
		<Tabs.Root value="sign-in" class="w-full">
			<Tabs.List class="w-full">
				<Tabs.Trigger value="sign-in">Sign In</Tabs.Trigger>
				<Tabs.Trigger value="register">Register</Tabs.Trigger>
			</Tabs.List>
			<Tabs.Content value="sign-in">
				<LoginForm form={loginForm} closeDialog={() => props.onOpenChange?.(false)} />
			</Tabs.Content>
			<Tabs.Content value="register"></Tabs.Content>
		</Tabs.Root>
	</Dialog.Content>
</Dialog.Root>
