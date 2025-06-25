
import axios from "axios";

import { UserClient } from "./components/client";
import { UserColumn } from "./components/columns";
import { cookies } from 'next/headers';



const UsersPage = async () => {
    let formattedUsers: UserColumn[] = [];
    const cookieStore = cookies();
    const token = cookieStore.get('token')?.value;


    try {
        const users = await axios.get(`${process.env.PV_PATH_API}/users?token=${token}`);

        formattedUsers = users.data.map((user: any) => ({
            id: user.u_id,
            email: user.u_email,
            uid: user.u_uid,
            registerDate: user.u_register_date,
            isDisabled: user.u_is_disabled,
            userStatus: user.user_status,
        }));

    } catch (error) {
        formattedUsers = [];
    }


    return (
        <div className="flex-col">
            <div className="flex-1 space-y-4 p-8 pt-6">
                <UserClient data={formattedUsers} />
            </div>
        </div>
    );
};

export default UsersPage;