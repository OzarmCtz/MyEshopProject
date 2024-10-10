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
import { toast } from "react-hot-toast";
import { useRouter } from 'next/navigation';
import { getAuth, createUserWithEmailAndPassword } from 'firebase/auth';
import Link from "next/link";

const formSchema = z.object({
    emailAddress: z.string().email(),
    password: z.string().min(6),
    confirmPassword: z.string().min(6)
}).refine(data => data.password === data.confirmPassword, {
    message: "Passwords don't match",
    path: ["confirmPassword"], // Path of error
});

export default function SignUpPage() {
    const [isLoading, setLoading] = useState(false);
    const router = useRouter();

    const form = useForm<z.infer<typeof formSchema>>({
        resolver: zodResolver(formSchema),
        defaultValues: {
            emailAddress: "",
            password: "",
            confirmPassword: "",
        },
    });

    const handleSignUp = async (values: z.infer<typeof formSchema>) => {
        try {
            setLoading(true);
            const auth = getAuth();
            const result = await createUserWithEmailAndPassword(auth, values.emailAddress, values.password);
            console.log("SignUpPage - Sign up result:", result);
            if (result.user) {
                const response = await fetch(`${process.env.NEXT_PUBLIC_PATH_API}/sign/up`, {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({
                        u_email: result.user.email,
                        u_uid: result.user.uid,
                        u_is_disabled: false // Assurez-vous que c'est un boolean
                    }),
                });

                const data = await response.json();
                console.log("Backend response:", data);

                if (!response.ok) {
                    throw new Error(`Failed to register user with backend: ${data.message || response.statusText}`);
                }

                toast.success("Account created successfully. Redirecting to sign in...");
                router.push('/sign-in');
            }
        } catch (error: any) {
            console.error("Sign up error:", error);
            toast.error(error.message || "An error occurred during sign up.");
        } finally {
            setLoading(false);
        }
    };

    return (
        <main className="flex min-h-screen flex-col items-center justify-between p-24">
            <Form {...form}>
                <form
                    onSubmit={form.handleSubmit(handleSignUp)}
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
                    <FormField
                        control={form.control}
                        name="confirmPassword"
                        render={({ field }) => (
                            <FormItem>
                                <FormLabel>Confirm Password</FormLabel>
                                <FormControl>
                                    <Input
                                        disabled={isLoading}
                                        {...field}
                                        placeholder="Confirm Password"
                                        type="password"
                                    />
                                </FormControl>
                                <FormMessage />
                            </FormItem>
                        )}
                    />
                    <Button type="submit" className="w-full" disabled={isLoading}>
                        {isLoading ? 'Loading...' : 'Sign Up'}
                    </Button>
                    <div className="mt-4 text-center">
                        <p>
                            Already have an account?{" "}
                            <Link href="/sign-in" className="text-blue-500">
                                Sign in
                            </Link>
                        </p>
                    </div>
                </form>
            </Form>
        </main>
    );
}
