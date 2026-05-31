'use server'

import { cookies } from 'next/headers'
import { createClient } from '@/utils/supabase/server'
import { redirect } from 'next/navigation'

export async function login(_prevState: any, formData: FormData) {
  // The form data is extracted
  const email = formData.get('email') as string
  const password = formData.get('password') as string

  // Are prepared the cookies and the client of Supabase
  const cookieStore = await cookies()
  const supabase = createClient(cookieStore)

  // The request is made by default
  const { error } = await supabase.auth.signInWithPassword({
    email,
    password,
  })

  // Handling of results
  if (error) {
    // If the password is incorrect or user don't exist, it redirects with an error parameter
    return {error: "Invalid email or password"}
  }

  // If everything is correct, the session is saved in cookies and redirected to the main page
  redirect('/mainMenu')
}

export async function signUp(_prevState: any, formData: FormData) {
  const email = formData.get('email') as string
  const password = formData.get('password') as string

  const cookieStore = await cookies()
  const supabase = createClient(cookieStore)

  const { error } = await supabase.auth.signUp({
    email,
    password,
  })

  if (error) {
    return {
      error: error.message,
    }
  }

  redirect('/mainMenu')
}