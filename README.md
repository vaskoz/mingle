# Mingle
For when you have an in-person offsite and want to make sure people meet each other.

# Running
```
MINGLE_TEAM_DIR=$HOME/offsiteTeams ./mingle 12 12 12 12 12 12 11
```

# The Office (US TV show) Offsite

A directory with all the different Dunder Mifflin branches. Inside each file is a list of employees at that branch.
```
Akron branch      Albany branch     Corporate HQ      Nashua branch
Rochester branch  Scranton branch   Stamford branch
```

Here is a view of the employees in each branch.
```
~/TheUsaOfficeOffsite> for f in *; echo $f; echo '-------------------';cat $f;echo;end
Akron branch
-------------------
Haunted Manager
Haunted Salesperson
Haunted Accountant

Albany branch
-------------------
Jeff

Corporate HQ
-------------------
Jo Bennett
David Wallace
Jan Levinson

Nashua branch
-------------------
Holly Flax

Rochester branch
-------------------
Mark Chisholm

Scranton branch
-------------------
Michael Scott
Jim Halpert
Pam Beesly
Dwight Schrute
Andy Bernard
Erin Hannon
Phyllis Vance
Stanley Hudson

Stamford branch
-------------------
Karen Filippelli
Josh Porter
Martin Nash
Hannah Smoterich-Barr
Tony Gardner
```

Finally, run the offsite mingling program.

```
MINGLE_TEAM_DIR=$HOME/TheUsaOfficeOffsite ./mingle 6 5 6 5
=====================================================
[GROUP 1] Seating Arrangement for 6 people.
=====================================================
Dwight Schrute from Scranton branch
Karen Filippelli from Stamford branch
Haunted Salesperson from Akron branch
Jan Levinson from Corporate HQ
Jeff from Albany branch
Holly Flax from Nashua branch
=====================================================
[GROUP 2] Seating Arrangement for 5 people.
=====================================================
Michael Scott from Scranton branch
Hannah Smoterich-Barr from Stamford branch
Haunted Accountant from Akron branch
David Wallace from Corporate HQ
Mark Chisholm from Rochester branch
=====================================================
[GROUP 3] Seating Arrangement for 6 people.
=====================================================
Andy Bernard from Scranton branch
Martin Nash from Stamford branch
Haunted Manager from Akron branch
Jo Bennett from Corporate HQ
Pam Beesly from Scranton branch
Tony Gardner from Stamford branch
=====================================================
[GROUP 4] Seating Arrangement for 5 people.
=====================================================
Erin Hannon from Scranton branch
Josh Porter from Stamford branch
Phyllis Vance from Scranton branch
Jim Halpert from Scranton branch
Stanley Hudson from Scranton branch
```
