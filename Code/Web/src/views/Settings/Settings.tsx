import React, { useEffect, useState } from 'react'
import { useTranslation } from 'react-i18next'
import i18n from 'i18next'
import {
  Button,
  Segmented,
  Upload,
  UploadFile,
  UploadProps,
  Input,
  message
} from 'antd'
import { LogoutOutlined } from '@ant-design/icons'

import { useAuth } from '@/context/authContext'
import SubtitledElement from '@/components/SubtitledElement/SubtitledElement'
import EditIcon from '@/assets/icons/edit.png'
import SaveIcon from '@/assets/icons/save.png'
import CloseIcon from '@/assets/icons/close.png'
import DefaultUser from '@/assets/images/DefaultProfile.png'
import UpdateUserInfos from '@/services/api/User/UpdateUserInfos'
import GetUserPicture from '@/services/api/User/GetUserPicture'
import base64ToFile from '@/utils/base64/baseToFile'
import PutUserPicture from '@/services/api/User/PutUserPicture'
import style from './Settings.module.css'

interface UserSettingsProps {
  t: (key: string) => string
}

const UserSettings: React.FC<UserSettingsProps> = ({ t }) => {
  const { user, updateUser } = useAuth()
  const [fileList, setFileList] = useState<UploadFile[]>([]);
  const [picture, setPicture] = useState<string | null>(null);
  const [isUploading, setIsUploading] = useState(false);

  const putUserPicture = async (pictureData: string) => {
    if (isUploading) return;

    setIsUploading(true);
    try {
      const req = await PutUserPicture(pictureData.split(',')[1]);
      if (req) {
        message.success(t('components.messages.picture_updated'));
      } else {
        message.error(t('components.messages.picture_not_updated'));
      }
    } catch (error) {
      console.error('Error updating user picture:', error);
    } finally {
      setIsUploading(false);
    }
  };

  const handleChange: UploadProps['onChange'] = ({ fileList: newFileList }) => {
    setFileList(newFileList);
  };

  const customRequest: UploadProps['customRequest'] = async ({ file }) => {
    try {
      const reader = new FileReader();
      reader.onload = (e) => {
        if (e.target?.result) {
          const base64 = e.target.result as string;
          setPicture(base64);
          putUserPicture(base64);
        }
      };
      reader.readAsDataURL(file as File);
    } catch (error) {
      message.error(t('components.messages.error_uploading_picture'));
    }
  };

  const beforeUpload = (file: File) => {
    const isImage = file.type.startsWith('image/');
    if (!isImage) {
      message.error('Only image files are allowed');
    }
    return isImage;
  };

  useEffect(() => {
    if (!user) return;
    const fetchPicture = async () => {
      try {
        const picture = await GetUserPicture(user?.id || '');
        const file = base64ToFile(picture.data, 'user.jpg', 'image/jpeg');
        const imageUrl = URL.createObjectURL(file);
        setPicture(imageUrl);
      } catch (error) {
        console.error('Error fetching user picture:', error);
      }
    };
    fetchPicture();
  }, [user?.id]);

  const [editData, setEditData] = useState(false)
  const [oldData, setOldData] = useState({
    firstname: user?.firstname,
    lastname: user?.lastname,
    email: user?.email
  })
  const [newData, setNewData] = useState({
    firstname: user?.firstname,
    lastname: user?.lastname,
    email: user?.email
  })

  const saveNewData = async () => {
    if (
      newData.firstname === oldData.firstname &&
      newData.lastname === oldData.lastname
    ) {
      setEditData(false);
      message.info(t('components.messages.no_modifications'));
      return;
    }
    try {
      await UpdateUserInfos({
        firstname: newData.firstname as string,
        lastname: newData.lastname as string,
      });
      setOldData(newData);
      updateUser(newData);
      setEditData(false);
      message.success(t('components.messages.modifications_saved'));
    } catch (error) {
      console.error('Error updating user data:', error);
    }
  };

  const cancelEdit = () => {
    setNewData(oldData);
    setEditData(false);
  };

  useEffect(() => {
    if (user) {
      setNewData({
        firstname: user.firstname,
        lastname: user.lastname,
        email: user.email,
      });
      setOldData({
        firstname: user.firstname,
        lastname: user.lastname,
        email: user.email,
      });
    }
  }, [user]);

  return (
    <div className={style.settingsContainer}>
      <div className={style.userItem}>
      <Upload
        fileList={fileList}
        onChange={handleChange}
        beforeUpload={beforeUpload}
        showUploadList={false}
        listType="picture-circle"
        maxCount={1}
        customRequest={customRequest}
      >
        <img
          src={picture || DefaultUser}
          alt="user"
          className={style.image}
        />
      </Upload>
      </div>
      <div className={style.userInformations}>
        <div className={style.titleContainer}>
          <b>{t('pages.settings.user_infos')}</b>
          <div className={style.editButtons}>
            {editData && (
              <Button
                type="link"
                style={{ width: 25, height: 25, padding: 10 }}
                onClick={cancelEdit}
              >
                <img
                  src={CloseIcon}
                  alt="edit"
                  style={{ width: 20, height: 20 }}
                />
              </Button>
            )}
            <Button
              type="link"
              style={{ width: 25, height: 25, padding: 10 }}
              onClick={() => editData ? saveNewData() : setEditData(!editData)}
            >
              <img
                src={editData ? SaveIcon : EditIcon}
                alt="edit"
                style={{ width: 20, height: 20 }}
              />
            </Button>
          </div>
        </div>
        <SubtitledElement subtitleKey={t('components.input.first_name.label')} subTitleStyle={{ opacity: 0.6 }}>
          {!editData ? (
            user?.firstname
          ) : (
            <Input
              defaultValue={user?.firstname}
              onChange={e =>
                setNewData({ ...newData, firstname: e.target.value })
              }
            />
          )}
        </SubtitledElement>
        <SubtitledElement subtitleKey={t('components.input.last_name.label')} subTitleStyle={{ opacity: 0.6 }}>
          {!editData ? (
            user?.lastname
          ) : (
            <Input
              defaultValue={user?.lastname}
              onChange={e =>
                setNewData({ ...newData, lastname: e.target.value })
              }
            />
          )}
        </SubtitledElement>
        <SubtitledElement subtitleKey={t('components.input.email.label')} subTitleStyle={{ opacity: 0.6 }}>
          {user?.email}
        </SubtitledElement>
      </div>
    </div>
  )
}

const Settings: React.FC = () => {
  const { t } = useTranslation()
  const { logout } = useAuth()

  const switchLanguage = (language: string) => {
    let lang = ''
    switch (language) {
      case 'fr' as string:
        lang = 'fr'
        localStorage.setItem('lang', 'fr')
        break
      case 'en' as string:
        lang = 'en'
        localStorage.setItem('lang', 'en')
        break
      default:
        lang = 'fr'
        localStorage.setItem('lang', 'fr')
        break
    }
    i18n.changeLanguage(lang)
  }

  return (
    <div className={style.layoutContainer}>
      <UserSettings t={t} />

      <div className={style.settingsContainer}>
        <div className={style.settingsItem}>
          {t('pages.settings.language')}
          <Segmented
            options={[
              { label: t('pages.settings.fr'), value: 'fr' },
              { label: t('pages.settings.en'), value: 'en' }
            ]}
            value={i18n.language}
            onChange={value => switchLanguage(value as string)}
          />
        </div>

        <div className={style.settingsItem}>
          {t('pages.settings.logout')}
          <Button
            type="primary"
            danger
            shape="circle"
            icon={<LogoutOutlined />}
            onClick={() => logout()}
          />
        </div>
      </div>
    </div>
  )
}

export default Settings
