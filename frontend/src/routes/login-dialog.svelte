<script lang="ts">
	import Button from '$lib/components/ui/button/button.svelte';
	import * as Dialog from '$lib/components/ui/dialog/index';
	import * as Tabs from '$lib/components/ui/tabs/index';
	import type { ComponentProps } from 'svelte';
	import LoginForm from './login-form.svelte';
	import RegisterForm from './register-form.svelte';

	type Props = Pick<ComponentProps<typeof Dialog.Root>, 'open' | 'onOpenChange'> & {
		loginForm: ComponentProps<typeof LoginForm>['form'];
		registerForm: ComponentProps<typeof RegisterForm>['form'];
	};

	const { loginForm, registerForm, ...props }: Props = $props();
	let selectedTab = $state('sign-in');
</script>

<Dialog.Root {...props}>
	<Dialog.Content>
		<Dialog.Header>
			<Dialog.Title>{selectedTab === 'sign-in' ? 'Sign In' : 'Register'}</Dialog.Title>
			<Dialog.Description
				>{selectedTab === 'sign-in'
					? 'Login to your account'
					: 'Register a new account'}</Dialog.Description
			>
		</Dialog.Header>
		<Tabs.Root value={selectedTab} onValueChange={(val) => (selectedTab = val)} class="w-full">
			<Tabs.List class="w-full">
				<Tabs.Trigger value="sign-in">Sign In</Tabs.Trigger>
				<Tabs.Trigger value="register">Register</Tabs.Trigger>
			</Tabs.List>
			<Tabs.Content value="sign-in">
				<LoginForm form={loginForm} closeDialog={() => props.onOpenChange?.(false)} />
			</Tabs.Content>
			<Tabs.Content value="register">
				<RegisterForm form={registerForm} closeDialog={() => props.onOpenChange?.(false)} />
			</Tabs.Content>
		</Tabs.Root>
	</Dialog.Content>
</Dialog.Root>
