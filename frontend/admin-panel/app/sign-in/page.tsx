"use client";

import { useState } from 'react';
import * as z from "zod";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import {
    Form,
    FormControl,
    FormField,
    FormItem,
    FormLabel,
    FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { toast, Toaster } from "react-hot-toast";
import Link from "next/link";

import { useLogin } from '@/hooks/useLogin';
import { useRouter } from 'next/navigation';

const formSchema = z.object({
    emailAddress: z.string().email(),
    password: z.string().min(6),
});

export default function SignInPage() {
    const [isLoading, setLoading] = useState(false);
    const router = useRouter();

    const form = useForm<z.infer<typeof formSchema>>({
        resolver: zodResolver(formSchema),
        defaultValues: {
            emailAddress: "",
            password: "",
        },
    });

    const { login } = useLogin();

    const handleSubmit = async (values: z.infer<typeof formSchema>) => {
        try {
            setLoading(true);
            const result = await login(values.emailAddress, values.password);
            if (result.user) {
                toast.success("Signed in successfully. Redirecting to dashboard...");
                router.push('/');
            }
        } catch (error: any) {
            toast.error(error.message || "An error occurred during sign in.");
        } finally {
            setLoading(false);
        }
    };

    return (
        <main className="flex min-h-screen flex-col items-center justify-between p-24">
            <Form {...form}>
                <form
                    onSubmit={form.handleSubmit(handleSubmit)}
                    className="max-w-md w-full flex flex-col gap-4"
                >
                    <FormField
                        control={form.control}
                        name="emailAddress"
                        render={({ field }) => (
                            <FormItem>
                                <FormLabel>Email address</FormLabel>
                                <FormControl>
                                    <Input
                                        disabled={isLoading}
                                        {...field}
                                        placeholder="Email address"
                                        type="email"
                                    />
                                </FormControl>
                                <FormMessage />
                            </FormItem>
                        )}
                    />
                    <FormField
                        control={form.control}
                        name="password"
                        render={({ field }) => (
                            <FormItem>
                                <FormLabel>Password</FormLabel>
                                <FormControl>
                                    <Input
                                        disabled={isLoading}
                                        {...field}
                                        placeholder="Password"
                                        type="password"
                                    />
                                </FormControl>
                                <FormMessage />
                            </FormItem>
                        )}
                    />
                    <Button type="submit" className="w-full" disabled={isLoading}>
                        {isLoading ? 'Loading...' : 'Submit'}
                    </Button>
                    {/* <div className="mt-4 text-center">
                        <p>
                            Don't have an account?{" "}
                            <Link href="/sign-up" className="text-blue-500">
                                Sign up
                            </Link>
                        </p>
                    </div> */}
                </form>
            </Form>
        </main>
    );
}
