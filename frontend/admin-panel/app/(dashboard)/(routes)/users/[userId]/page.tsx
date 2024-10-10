import { User } from '@/app/schema/schema';
import { UserForm } from './components/user-form';
import axios from 'axios';
import { cookies } from 'next/headers';

const ItemPage = async ({ params }: { params: { userId: number } }) => {
    let userData: User | null = null;



    const cookieStore = cookies();
    const token = cookieStore.get('token')?.value;
    if (params.userId) {
        try {
            const response = await axios.get(`${process.env.NEXT_PUBLIC_PV_PATH_API}/users/${params.userId}?token=${token}`);
            if (response.status === 200) {
                userData = response.data;
            }
        } catch (error) {
            console.error("Failed to fetch item:", error);
        }

    }

    return (
        <div className="flex-col">
            <div className="flex-1 space-y-4 p-8 pt-6">
                <UserForm
                    initialData={userData}
                />
            </div>
        </div>
    );
}

export default ItemPage;
