package repository

// pengecekan permission berdarsarkan id dan code nya
func (r *Repository) Allowed(userID uint, code string) (bool, error) {
	var exists bool

	query := `
        SELECT EXISTS (
            SELECT 1
            FROM users u
            JOIN role_permissions rp ON rp.role_id = u.role_id
            JOIN permissions p ON p.id = rp.permission_id
            WHERE u.id = ? AND p.code = ?
        )
    `
	err := r.db.Raw(query, userID, code).Scan(&exists).Error

	if err != nil {
		return false, err
	}

	return exists, nil
}

// menemukan id user dari session Token
func (r *Repository) FindUserIDBySession(sessionToken string) (uint, error) {
	var userID uint
	err := r.db.Select("user_id").Where("token=?", sessionToken).First(&userID).Error
	if err != nil {
		return 0, err
	}
	return userID, nil
}
