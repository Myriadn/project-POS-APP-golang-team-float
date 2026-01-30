package repository

// pengecekan permission berdarsarkan id dan code nya
func (r *Repository) Allowed(userID uint, code string) (bool, error) {
	var exists bool

	query := `
        WITH perm AS (
        SELECT id FROM permissions WHERE code = ?
    )
    SELECT
    CASE
        WHEN EXISTS (
            SELECT 1 FROM user_permissions up, perm
            WHERE up.user_id = ?
            AND up.permission_id = perm.id 
        ) THEN FALSE
        
 
        WHEN EXISTS (
            SELECT 1 FROM user_permissions up, perm
            WHERE up.user_id = ? 
            AND up.permission_id = perm.id 
        ) THEN TRUE
        
        WHEN EXISTS (
            SELECT 1
            FROM users u
            JOIN role_permissions rp ON rp.role_id = u.role_id
            JOIN perm ON perm.id = rp.permission_id
            WHERE u.id = ?
        ) THEN TRUE
        
        ELSE FALSE
    END;
    `
	err := r.db.Raw(query, code, userID, userID, userID).Scan(&exists).Error

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
